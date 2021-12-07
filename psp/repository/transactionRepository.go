package repository

import (
	"context"
	"ubiquitous-payment/psp/model"
)

const transactionsCollectionName = "psp-transactions"

var emptyContext = context.TODO()

func (repo *Repository) CreateTransaction(transaction *model.Transaction) error {
	transactionsCollectionName := repo.getCollection(transactionsCollectionName)
	_, err := transactionsCollectionName.InsertOne(emptyContext, transaction)
	return err
}
