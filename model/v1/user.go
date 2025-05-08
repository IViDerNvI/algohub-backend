package model

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/algohub/pkg/util/bcryptutil"
	"github.com/ividernvi/algohub/pkg/util/idutil"

	"gorm.io/gorm"
)

type User struct {
	ObjMeta `json:",inline"`

	UserName string `json:"username" gorm:"column:username;unique;not null" validate:"required,min=8,max=16"`
	Password string `json:"password,omitempty" gorm:"column:password;not null" validate:"required,min=8,max=16"`

	Status string `json:"status" gorm:"column:status;not null" validate:"required,oneof=admin user guest"`

	NickName string `json:"nickname" gorm:"column:nickname;not null"`
	Email    string `json:"email" gorm:"column:email;not null" validate:"required,email"`
	Phone    string `json:"phone" gorm:"column:phone"`

	Avatar     string `json:"avatar" gorm:"column:avatar"`
	Bio        string `json:"bio" gorm:"column:bio"`
	Company    string `json:"company" gorm:"column:company"`
	Location   string `json:"location" gorm:"column:location"`
	ProfileURL string `json:"profile_url" gorm:"profile_url"`

	IpAddress string `json:"ip_address" gorm:"column:ip_address"`
}

type UserList struct {
	ListMeta `json:",inline"`
	Items    []User
}

func (u *User) IsAdmin() bool {
	return u.Status == "admin"
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	user_info, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(user_info)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.InstanceID = uint(idutil.SnowflakeID())
	var err error
	u.Password, err = bcryptutil.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) error {
	return nil
}

func (u *User) Validate() error {
	validator := validator.New()
	if err := validator.Struct(u); err != nil {
		return err
	}
	return nil
}

func (u *User) AfterGet(tx *gorm.DB) error {
	u.ID = 0
	u.Password = ""
	return nil
}

func (u *User) Override(newuser *User) *User {
	if newuser.Password != "" {
		var err error
		u.Password, err = bcryptutil.HashPassword(u.Password)
		if err != nil {
			return nil
		}
	}

	if newuser.NickName != "" {
		u.NickName = newuser.NickName
	}
	if newuser.Email != "" {
		u.Email = newuser.Email
	}
	if newuser.Phone != "" {
		u.Phone = newuser.Phone
	}
	if newuser.Avatar != "" {
		u.Avatar = newuser.Avatar
	}
	if newuser.Bio != "" {
		u.Bio = newuser.Bio
	}
	if newuser.Company != "" {
		u.Company = newuser.Company
	}
	if newuser.Location != "" {
		u.Location = newuser.Location
	}
	return u
}
