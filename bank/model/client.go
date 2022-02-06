package model

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	Username  				string          `json:"username" gorm:"not null;unique"`
	Password  				string          `json:"password" gorm:"not null"`
	IsDeleted 				bool            `json:"isDeleted" gorm:"not null"`
	Roles     				[]Role          `json:"roles" gorm:"many2many:user_roles;"`
	Accounts  				[]ClientAccount `json:"accounts" gorm:"many2many:user_accounts;"`
	LastActivityTimestamp	time.Time		`json:"lastActivityTimestamp"`
}
