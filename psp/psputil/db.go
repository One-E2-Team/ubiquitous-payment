package psputil

import "context"

const (
	PspDbName                  = "psp-db"
	WebShopCollectionName      = "psp-clients"
	TransactionsCollectionName = "psp-transactions"
	PaymentTypesCollectionName = "paymentTypes"
	AccountsCollectionName     = "accounts"

	PSPIDFieldName = "pspid"
	IDFieldName    = "_id"
)

var EmptyContext = context.TODO()
