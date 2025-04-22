package model

import (
	"strings"

	"gorm.io/gorm"
)

type OptMeta struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type CreateOptions struct {
	ObjMeta `json:",inline"`
	DryRun  bool `json:"dryRun"`
}

type DeleteOptions struct {
	ObjMeta            `json:",inline"`
	DryRun             bool  `json:"dryRun"`
	GracePeriodSeconds int64 `json:"gracePeriodSeconds"`
}

type GetOptions struct {
	ObjMeta `json:",inline"`
	DryRun  bool `json:"dryRun"`
}

type ListOptions struct {
	ObjMeta         `json:",inline"`
	DryRun          bool   `json:"dryRun"`
	Limit           int    `json:"limit" validate:"min=1,max=100,required"`
	Offset          int    `json:"offset" validate:"min=0,required"`
	Selector        string `json:"selector" validate:"required"`
	ResourceVersion string `json:"resourceVersion"`
	TimeoutSeconds  int    `json:"timeoutSeconds"`
}

type UpdateOptions struct {
	ObjMeta `json:",inline"`
	DryRun  bool `json:"dryRun"`
}

type VerifyOptions struct {
	ObjMeta  `json:",inline"`
	DryRun   bool `json:"dryRun"`
	IsBasic  bool `json:"isBasic"`
	IsBearer bool `json:"isBearer"`
}

// GetSearchKey returns the search key for the ListOptions.
// It is used to filter the list of resources based on the search key.
//
// Selector Example:
//
//	{
//		...
//		"selector": "username~admin,status=admin,email!xxx@email.com",
//		...
//	}
//
// Equal to:
//
//	SELECT * FROM TABLE WHERE username LIKE admin AND status = admin AND email != xxx.@email.com
func (opts *ListOptions) ApplyListOptions(db *gorm.DB) *gorm.DB {
	if opts.Selector == "" {
		return db
	}

	conditions := strings.Split(opts.Selector, ",")
	for _, condition := range conditions {
		if strings.Contains(condition, "=") {
			parts := strings.Split(condition, "=")
			db = db.Where(parts[0]+" LIKE ?", parts[1])
		} else if strings.Contains(condition, "!") {
			parts := strings.Split(condition, "!")
			db = db.Where(parts[0]+" != ?", parts[1])
		} else if strings.Contains(condition, "~") {
			parts := strings.Split(condition, "~")
			db = db.Where(parts[0]+" LIKE ?", "%"+parts[1]+"%")
		}
	}

	return db
}

func Selector(searchkey map[string]string) string {
	var selector string
	for k, v := range searchkey {
		selector += k + "=" + "%" + v + "%,"
	}
	return selector[:len(selector)-1]
}

func (opts *ListOptions) Complete() {
	if opts.Limit == 0 {
		opts.Limit = 10
	}
	if opts.Offset == 0 {
		opts.Offset = 0
	}
	if opts.TimeoutSeconds == 0 {
		opts.TimeoutSeconds = 30
	}
	if opts.ResourceVersion == "" {
		opts.ResourceVersion = "v1"
	}
}
