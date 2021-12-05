package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ProfileId   	 uint           `json:"profileId" gorm:"not null"`
	Email 			 string 		`json:"email" gorm:"unique;not null"`
	Username		 string		    `json:"username" gorm:"not null;unique"`
	Password 		 string 		`json:"password" gorm:"not null"`
	IsDeleted 		 bool		    `json:"isDeleted" gorm:"not null"`
	IsValidated		 bool		    `json:"isValidated" gorm:"not null"`
	ValidationUuid   string         `json:"validationUuid"`
	ValidationExpire time.Time      `json:"validationExpire"`
	Roles            []Role         `json:"roles" gorm:"many2many:user_roles;"`
}
