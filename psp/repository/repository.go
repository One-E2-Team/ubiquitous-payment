package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ubiquitous-payment/psp/psputil"
)

type Repository struct {
	Client *mongo.Client
}

func (repo *Repository) getCollection(collectionName string) *mongo.Collection {
	return repo.Client.Database(psputil.PspDbName).Collection(collectionName)
}
