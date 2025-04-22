package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/algohub/pkg/util/idutil"
	"gorm.io/gorm"
)

type Solution struct {
	ObjMeta `json:",inline"`

	ProblemID  string `json:"problem_id" gorm:"column:problem_id" validate:"required"`
	TestData   string `json:"data_test" gorm:"column:data_test" validate:"required"`
	TestResult string `json:"result_test" gorm:"column:result_test" validate:"required"`
	Provider   string `json:"provider" gorm:"column:provider" validate:"required"`
}

type SolutionList struct {
	ListMeta `json:",inline"`
	Items    []Solution `json:"items"`
}

func (s *Solution) TableName() string {
	return "solutions"
}

func (s *Solution) BeforeCreate(tx *gorm.DB) error {
	s.InstanceID = uint(idutil.SnowflakeID())
	return nil
}

func (s *Solution) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Solution) Validate() error {
	validator := validator.New()
	if err := validator.Struct(s); err != nil {
		return err
	}
	return nil
}

func (s *Solution) Override(new *Solution) *Solution {
	s.TestData = new.TestData
	s.TestResult = new.TestResult
	s.Provider = new.Provider
	return s
}
