package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"ubiquitous-payment/psp/model"
)

const transactionsCollectionName = "psp-transactions"

var emptyContext = context.TODO()

func (repo *Repository) CreateTransaction(transaction *model.Transaction) error {
	transactionsCollection := repo.getCollection(transactionsCollectionName)
	_, err := transactionsCollection.InsertOne(emptyContext, transaction)
	return err
}

func (repo *Repository) GetTransactionByPspId(pspId string) (*model.Transaction,error) {
	transactionsCollection := repo.getCollection(transactionsCollectionName)
	filter := bson.D{{"pspid", pspId}}
	var result model.Transaction
	err := transactionsCollection.FindOne(emptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) UpdateTransaction(transaction *model.Transaction) error {
	transactionsCollection := repo.getCollection(transactionsCollectionName)
	filter := bson.D{{"pspid", transaction.PSPId}}
	_, err := transactionsCollection.ReplaceOne(emptyContext,filter,transaction)
	return err
}
