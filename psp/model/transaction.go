package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"ubiquitous-payment/psp/psputil"
)

type Transaction struct {
	ID                    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	PSPId                 string             `json:"pspId"`
	WebShopID             string             `json:"webShopID"`
	Amount                float32            `json:"amount"`
	Currency              string             `json:"Currency"`
	SuccessURL            string             `json:"successURL"`
	FailURL               string             `json:"failURL"`
	ErrorURL              string             `json:"errorURL"`
	MerchantOrderID       string             `json:"merchantOrderID"`
	MerchantTimestamp     time.Time          `json:"merchantTimestamp"`
	PaymentMode           PaymentMode        `json:"paymentMode"`
	Recurring             *Recurring         `json:"recurring"`
	IsSubscription        bool               `json:"isSubscription"`
	TransactionStatus     TransactionStatus  `json:"transactionStatus"`
	AvailablePaymentTypes []PaymentType      `json:"availablePaymentTypes"`
	SelectedPaymentType   PaymentType        `json:"selectedPaymentType"`
	MerchantAccounts      []Account          `json:"merchant_accounts"`
	ExternalTransactionId string             `json:"externalTransactionId"`
}

func (transaction *Transaction) GetURLByStatus() string {
	switch transaction.TransactionStatus {
	case FULFILLED:
		return transaction.SuccessURL
	default:
		return transaction.FailURL
	}
}

func (transaction *Transaction) IsPricingPlan(plugin psputil.Plugin) (bool, error) {
	if (transaction.PaymentMode == ONE_TIME && transaction.IsSubscription) || (transaction.PaymentMode == RECURRING) {
		if !plugin.SupportsPlanPayment() {
			return false, errors.New("plugin does not support pricing plan")
		}
		return true, nil
	}
	return false, nil
}

func (transaction *Transaction) GetSelectedAccount() (Account, error) {
	for _, acc := range transaction.MerchantAccounts {
		if acc.PaymentType.Name == transaction.SelectedPaymentType.Name {
			return acc, nil
		}
	}
	return Account{}, errors.New("don't have selected account")
}
