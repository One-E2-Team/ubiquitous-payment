package model

import "gorm.io/gorm"

type WebShop struct {
	gorm.Model
	Name           string `json:"name" gorm:"unique;not null"`
	PSPAccessToken string `json:"pspAccessToken"`
}
