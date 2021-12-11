package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) CreateWebShop(webShop *model.WebShop) error {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	_, err := webShopCollection.InsertOne(psputil.EmptyContext, webShop)
	return err
}

func (repo *Repository) ChangeWebShopAcceptance(webShopID primitive.ObjectID, isAccepted bool) error {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	updateFilter := bson.D{{psputil.SetSelector, bson.D{{psputil.AcceptedFieldName, isAccepted}}}}
	_, err := webShopCollection.UpdateByID(psputil.EmptyContext, webShopID, updateFilter)
	return err
}
