package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) GetWebShopByID(webShopID primitive.ObjectID) (*model.WebShop, error) {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	filter := bson.D{{psputil.IDFieldName, webShopID}}
	var result model.WebShop
	err := webShopCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) GetWebShopByIDString(webShopID string) (*model.WebShop, error) {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	filter := bson.D{{psputil.IDFieldName, webShopID}}
	var result model.WebShop
	err := webShopCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) GetWebShopByName(webShopName string) (*model.WebShop, error) {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	filter := bson.D{{psputil.NameFieldName, webShopName}}
	var result model.WebShop
	err := webShopCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) ChangeWebShopAcceptance(webShopID primitive.ObjectID, isAccepted bool) error {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	updateFilter := bson.D{{psputil.SetSelector, bson.D{{psputil.AcceptedFieldName, isAccepted}}}}
	_, err := webShopCollection.UpdateByID(psputil.EmptyContext, webShopID, updateFilter)
	return err
}

func (repo *Repository) UpdateWebShop(webShop *model.WebShop) error {
	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	filter := bson.D{{psputil.IDFieldName, webShop.ID}}
	_, err := webShopCollection.ReplaceOne(psputil.EmptyContext, filter, webShop)
	return err
}
