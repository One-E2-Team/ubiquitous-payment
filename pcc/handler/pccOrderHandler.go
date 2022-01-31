package handler

import (
	"net/http"
	"ubiquitous-payment/pcc/dto"
	"ubiquitous-payment/pcc/handler/mapper"
	"ubiquitous-payment/util"
)

func (handler *Handler) CreatePccOrder(w http.ResponseWriter, r *http.Request){
	var pccOrderDto dto.PccOrderDto
	err := util.UnmarshalRequest(r, &pccOrderDto)
	if err != nil{
		util.HandleErrorInHandler(err, w, loggingClass+"CreatePccOrder", loggingService)
		return
	}

	pccOrder := mapper.PccOrderDtoToPccOrder(pccOrderDto)

	err = handler.Service.CreatePccOrder(pccOrder)
	if err != nil{
		util.HandleErrorInHandler(err, w, loggingClass+"CreatePccOrder", loggingService)
		return
	}
}

