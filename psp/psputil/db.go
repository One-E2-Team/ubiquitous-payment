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

	PSPIDFieldName     = "pspid"
	IDFieldName        = "_id"
	NameFieldName      = "name"
	AcceptedFieldName  = "accepted"
	UsernameFieldName  = "username"
	WebShopIDFieldName = "webshopid"

	WebShopTokenPermissionName = "READ_ACCESS_TOKEN"
	WebShopOrderPermissionName = "CREATE_ORDER_FROM_WEB_SHOP"
	WebShopRoleName            = "WebShop"
)

var EmptyContext = context.TODO()
