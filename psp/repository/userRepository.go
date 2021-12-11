package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) CreateUser(user *model.User) error { //TODO: make only one method for create
	usersCollection := repo.getCollection(psputil.UsersCollectionName)
	_, err := usersCollection.InsertOne(psputil.EmptyContext, user)
	return err
}

func (repo *Repository) GetUserByUsername(username string) (*model.User, error) {
	usersCollection := repo.getCollection(psputil.UsersCollectionName)
	filter := bson.D{{psputil.UsernameFieldName, username}}
	var result model.User
	err := usersCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}
