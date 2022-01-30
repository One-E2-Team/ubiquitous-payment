package repository

import "gorm.io/gorm"

type Repository struct {
	Database *gorm.DB
}
