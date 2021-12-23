package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

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
	id, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{psputil.IDFieldName, id}}
	var transaction model.Transaction
	err = transactionsCollection.FindOne(psputil.EmptyContext, filter).Decode(&transaction)
	return transaction.AvailablePaymentTypes, err
}

func (repo *Repository) GetTransactionById(id primitive.ObjectID) (*model.Transaction, error) {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	filter := bson.D{{psputil.IDFieldName, id}}
	var result model.Transaction
	err := transactionsCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}

func (repo *Repository) GetTransactionByExternalId(externalId string) (*model.Transaction, error) {
	transactionsCollection := repo.getCollection(psputil.TransactionsCollectionName)
	filter := bson.D{{"externaltransactionid", externalId}}
	var result model.Transaction
	err := transactionsCollection.FindOne(psputil.EmptyContext, filter).Decode(&result)
	return &result, err
}
