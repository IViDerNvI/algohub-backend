package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/algohub/pkg/util/idutil"
	"gorm.io/gorm"
)

type Submit struct {
	ObjMeta `json:",inline"`

	CodeText       string  `json:"code_text" gorm:"column:code_text" validate:"required"`
	Language       string  `json:"language" gorm:"column:language" validate:"required"`
	ProblemID      string  `json:"problem_id" gorm:"column:problem_id" validate:"required"`
	Status         string  `json:"status" gorm:"column:status" validate:"required"`
	Author         string  `json:"author" gorm:"column:author" validate:"required"`
	ExcuteTime     int64   `json:"excute_time" gorm:"column:excute_time"`
	ExcuteMemory   float64 `json:"excute_memory" gorm:"column:excute_memory"`
	Details        string  `json:"details,omitempty" gorm:"column:details"`
	Input          string  `json:"input" gorm:"column:input"`
	ExpectedOutput string  `json:"expected_output" gorm:"column:expected_output"`
	ActualOutput   string  `json:"actual_output" gorm:"column:actual_output"`
}

var (
	SubmitStatusPending       = "pending"
	SubmitStatusAccepted      = "accepted"
	SubmitStatusWrong         = "outofmemory"
	SubmitStatusRuntime       = "runtimeerror"
	SubmitStatusCompile       = "compileerror"
	SubmitStatusTimeLimit     = "timelimit"
	SubmitStatusMemoryLimit   = "memorylimit"
	SubmitStatusSystemError   = "systemerror"
	SubmitStatusInternalError = "internalerror"
	SubmitStatusUnknown       = "unknown"
)

type SubmitList struct {
	ListMeta `json:",inline"`
	Items    []Submit `json:"items"`
}

func (s *Submit) TableName() string {
	return "submits"
}

func (s *Submit) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Submit) BeforeCreate(tx *gorm.DB) error {
	s.InstanceID = uint(idutil.SnowflakeID())
	return nil
}

func (s *Submit) Override(newSubmit *Submit) {
	if newSubmit.Status != "" {
		s.Status = newSubmit.Status
	}

	if newSubmit.Details != "" {
		s.Details = newSubmit.Details
	}
}
