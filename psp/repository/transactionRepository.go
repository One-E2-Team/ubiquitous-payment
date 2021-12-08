package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) CreateTransaction(transaction *model.Transaction) error {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	_, err := transactionsCollection.InsertOne(psputil.EmptyContext, transaction)
	return err
}

func (repo *Repository) GetTransactionByPspId(pspId string) (*model.Transaction, error) {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	filter := bson.D{{psputil.PspIdFieldName, pspId}}
	var result model.Transaction
	err := transactionsCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) UpdateTransaction(transaction *model.Transaction) error {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	filter := bson.D{{psputil.PspIdFieldName, transaction.PSPId}}
	_, err := transactionsCollection.ReplaceOne(psputil.EmptyContext, filter, transaction)
	return err
}
