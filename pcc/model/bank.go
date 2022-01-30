package model

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	PANPrefix     string `json:"panPrefix" gorm:"not null"`
	URL           string `json:"url"`
	IsActive 	  bool   `json:"isActive"`
}
