package repository

import (
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func (repo *Repository) Create(domainObject interface{}) error {
	result := repo.Database.Create(domainObject)
	if result.RowsAffected == 0 {
		return fmt.Errorf("object created")
	}
	return nil
}

func (repo *Repository) Update(domainObject interface{}) error {
	return repo.Database.Save(domainObject).Error
}
