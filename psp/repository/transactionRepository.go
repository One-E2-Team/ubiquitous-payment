package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	filter := bson.D{{psputil.PSPIDFieldName, pspId}}
	var result model.Transaction
	err := transactionsCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) UpdateTransaction(transaction *model.Transaction) error {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	filter := bson.D{{psputil.PSPIDFieldName, transaction.PSPId}}
	_, err := transactionsCollection.ReplaceOne(psputil.EmptyContext, filter, transaction)
	return err
}

func (repo *Repository) GetAvailablePaymentTypes(transactionID string) ([]model.PaymentType, error) {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	filter := bson.D{{psputil.IDFieldName, transactionID}}
	var transaction model.Transaction
	err := transactionsCollection.FindOne(psputil.EmptyContext, filter).Decode(&transaction)
	return transaction.AvailablePaymentTypes, err
}

func (repo *Repository) GetTransactionById(id primitive.ObjectID) (*model.Transaction, error) {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	filter := bson.D{{psputil.IDFieldName, id}}
	var result model.Transaction
	err := transactionsCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}