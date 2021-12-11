package repository

import (
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) CreateWebShop(webShop *model.WebShop) error {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	_, err := webShopCollection.InsertOne(psputil.EmptyContext, webShop)
	return err
}
