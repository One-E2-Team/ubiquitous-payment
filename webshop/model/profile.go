package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
}
