package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func (repo *Repository) GetAllBanks() ([]model.Bank, error) {
	banksCollection := repo.getCollection(psputil.BanksCollectionName)
	var ret []model.Bank
	cursor, err := banksCollection.Find(psputil.EmptyContext, bson.M{})
	for cursor.Next(psputil.EmptyContext) {
		var pt model.Bank
		if err = cursor.Decode(&pt); err != nil {
			return nil, err
		}
		ret = append(ret, pt)
	}
	return ret, err
}

func (repo *Repository) GetAllBanksKeyValue() (*map[string]string, error) {
	banks, err := repo.GetAllBanks()
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string, 0)
	for _, b := range banks {
		ret[b.PANPrefix] = b.URL
	}
	return &ret, nil
}
