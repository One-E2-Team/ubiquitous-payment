package psputil

import "context"

const (
	SetSelector = "$set"

	PspDbName                  = "psp-db"
	WebShopCollectionName      = "psp-clients"
	TransactionsCollectionName = "psp-transactions"
	PaymentTypesCollectionName = "paymentTypes"
	AccountsCollectionName     = "accounts"
	UsersCollectionName        = "users"

	PSPIDFieldName    = "pspid"
	IDFieldName       = "_id"
	NameFieldName     = "name"
	AcceptedFieldName = "accepted"

	WebShopTokenPermissionName = "READ_ACCESS_TOKEN"
	WebShopRoleName            = "WebShop"
)

var EmptyContext = context.TODO()
