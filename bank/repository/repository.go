package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func (repo *Repository) Update(domainObject interface{}) error {
	return repo.Database.Save(domainObject).Error
}
