package account

import (
	"netbelge/functions"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	Path        string `json:"path" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:text"`
	ParentID    uint   `json:"parent_id" gorm:"type:int;default:0"`
}

type Departments []Department

func (d *Department) BeforeCreate(tx *gorm.DB) (err error) {
	d.Path = functions.NormalizePath(d.Name)
	return
}

func (*Department) TableName() string {
	return "departments"
}
