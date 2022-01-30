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
	BanksCollectionName        = "banks"

	PSPIDFieldName     = "pspid"
	IDFieldName        = "_id"
	NameFieldName      = "name"
	AcceptedFieldName  = "accepted"
	UsernameFieldName  = "username"
	WebShopIDFieldName = "webshopid"
	PANPrefixFieldName = "panPrefix"

	WebShopTokenPrivilegeName = "READ_ACCESS_TOKEN"
	WebShopOrderPrivilegeName = "CREATE_ORDER_FROM_WEB_SHOP"
	WebShopRoleName           = "WebShop"
)

var EmptyContext = context.TODO()
