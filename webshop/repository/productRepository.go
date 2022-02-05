package repository

import (
	"errors"
	"ubiquitous-payment/webshop/model"
)

func (repo *Repository) CreateProduct(product *model.Product) error {
	result := repo.RelationalDatabase.Create(product)
	if result.RowsAffected == 0 {
		return errors.New("product was not created")
	}
	return nil
}

func (repo *Repository) GetProduct(productID uint) (*model.Product, error) {
	product := &model.Product{}
	if err := repo.RelationalDatabase.First(&product, "id = ?", productID).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *Repository) UpdateProduct(product *model.Product) error {
	return repo.RelationalDatabase.Save(product).Error
}

func (repo *Repository) GetActiveProducts() ([]model.Product, error) {
	var products []model.Product
	result := repo.RelationalDatabase.Table("products").Select("*").Where("is_active=1").Find(&products)
	return products, result.Error
}

func (repo *Repository) GetProductsByMerchantId(merchantId uint) ([]model.Product, error) {
	var products []model.Product
	if err := repo.RelationalDatabase.Find(&products, "merchant_profile_id = ?", merchantId).Error; err != nil {
		return nil, err
	}
	return products, nil
}
