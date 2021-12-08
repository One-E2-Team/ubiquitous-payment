package psputil

import "context"

const (
	PspDbName                  = "psp-db"
	WebShopCollectionName      = "psp-clients"
	TransactionsCollectionName = "psp-transactions"
	PaymentTypesCollectionName = "paymentTypes"
	AccountsCollectionName     = "accounts"

	PspIdFieldName = "pspid"
)

var EmptyContext = context.TODO()
