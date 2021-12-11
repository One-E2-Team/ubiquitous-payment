package repository

import (
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) CreateUser(user *model.User) error { //TODO: make only one method for create
	usersCollection := repo.getCollection(psputil.UsersCollectionName)
	_, err := usersCollection.InsertOne(psputil.EmptyContext, user)
	return err
}
