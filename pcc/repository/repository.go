package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	RelationalDatabase *gorm.DB
}

func (repo *Repository) Update(domainObject interface{}) error {
	return repo.RelationalDatabase.Save(domainObject).Error
}