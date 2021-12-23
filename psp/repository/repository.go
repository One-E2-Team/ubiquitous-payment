package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ubiquitous-payment/psp/psputil"
)

type Repository struct {
	Client *mongo.Client
}

func (repo *Repository) AddDBConstraints() error {
	usersCollection := repo.getCollection(psputil.UsersCollectionName)
	_, err := usersCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: psputil.UsernameFieldName, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return err
	}

	webShopCollection := repo.getCollection(psputil.WebShopCollectionName)
	_, err = webShopCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: psputil.NameFieldName, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return err
}

func (repo *Repository) Create(newEntity interface{}, collectionName string) error {
	collection := repo.getCollection(collectionName)
	_, err := collection.InsertOne(psputil.EmptyContext, newEntity)
	return err
}

func (repo *Repository) getCollection(collectionName string) *mongo.Collection {
	return repo.Client.Database(psputil.PspDbName).Collection(collectionName)
}
