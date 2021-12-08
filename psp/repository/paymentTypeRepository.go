package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) GetAllPaymentTypes() ([]model.PaymentType, error) {
	paymentTypesCollection := repo.getCollection(psputil.PaymentTypesCollectionName)
	var ret []model.PaymentType
	cursor, err := paymentTypesCollection.Find(psputil.EmptyContext, bson.M{})
	for cursor.Next(psputil.EmptyContext) {
		var pt model.PaymentType
		if err = cursor.Decode(&pt); err != nil {
			return nil, err
		}
		ret = append(ret, pt)
	}
	return ret, err
}
