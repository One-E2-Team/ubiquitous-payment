package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) GetUserByID(userID primitive.ObjectID) (*model.User, error) {
	usersCollection := repo.getCollection(psputil.UsersCollectionName)
	filter := bson.D{{psputil.IDFieldName, userID}}
	var result model.User
	err := usersCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) GetUserByUsername(username string) (*model.User, error) {
	usersCollection := repo.getCollection(psputil.UsersCollectionName)
	filter := bson.D{{psputil.UsernameFieldName, username}}
	var result model.User
	err := usersCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) GetUserByWebShopID(webShopID string) (*model.User, error) {
	fmt.Println(webShopID)
	usersCollection := repo.getCollection(psputil.UsersCollectionName)
	filter := bson.D{{psputil.WebShopIDFieldName, webShopID}}
	var result model.User
	err := usersCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}
