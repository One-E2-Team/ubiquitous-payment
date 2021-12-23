package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	RelationalDatabase *gorm.DB
}
