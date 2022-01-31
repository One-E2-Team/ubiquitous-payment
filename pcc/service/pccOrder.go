package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"ubiquitous-payment/pcc/dto"
	"ubiquitous-payment/pcc/model"
	"ubiquitous-payment/util"
)

func (service *Service) CreatePccOrder(pccOrder *model.PccOrder) error{
	pccOrder.IssuerTimestamp = time.Now()
	pccOrder.OrderStatus = model.PLACED
	err := service.Repository.CreatePccOrder(pccOrder)
	if err != nil{
		return err
	}
	return service.forwardOrderToIssuersBank(pccOrder)
}

func (service *Service) forwardOrderToIssuersBank(pccOrder *model.PccOrder) error{
	issuerPanPrefix := pccOrder.IssuerPAN[0:3]
	issuersBank, err := service.GetBankByPanPrefix(issuerPanPrefix)

	if err != nil{
		return err
	}

	client := &http.Client{}
	jsonReq, _ := json.Marshal(pccOrder)
	req, err := http.NewRequest(http.MethodPost, issuersBank.URL + "/pcc-issuer-pay", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp, err :=  client.Do(req)
	var respDto dto.IssuerBankResponseDto
	err = util.UnmarshalResponse(resp, &respDto)
	if err != nil{
		return err
	}
	pccOrder.IssuerOrderId = respDto.IssuerOrderId
	pccOrder.IssuerTimestamp = respDto.IssuerTimestamp
	pccOrder.OrderStatus = respDto.OrderStatus

	return service.Repository.Update(pccOrder)
}