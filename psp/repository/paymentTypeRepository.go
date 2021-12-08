package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"ubiquitous-payment/psp/model"
)

const paymentTypesCollectionName = "paymentTypes"

func (repo *Repository) GetAllPaymentTypes() ([]model.PaymentType,error) {
	transactionsCollection := repo.getCollection(paymentTypesCollectionName)
	var ret []model.PaymentType
	cursor, err := transactionsCollection.Find(emptyContext, bson.M{})
	for cursor.Next(emptyContext) {
		var pt model.PaymentType
		if err = cursor.Decode(&pt); err != nil {
			return nil, err
		}
		ret = append(ret, pt)
	}
	return ret, err
}
