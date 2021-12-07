package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Client *mongo.Client
}

const pspDbName = "psp-db"

func (repo *Repository) getCollection(collectionName string) *mongo.Collection {
	return repo.Client.Database(pspDbName).Collection(collectionName)
}
