package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/handler/mapper"
)

func (handler *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating product handler")
	if err := r.ParseMultipartForm(0); err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CreateProduct", loggingService)
		return
	}
	var request dto.ProductDTO
	data := r.MultipartForm.Value["data"]
	err := json.Unmarshal([]byte(data[0]), &request)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CreateProduct", loggingService)
		return
	}
	err = handler.WSService.CreateProduct(mapper.ProductDTOToProduct(request, util.GetLoggedUserIDFromToken(r)), r)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CreateProduct", loggingService)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (handler *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"UpdateProduct", loggingService)
		return
	}

	err = handler.WSService.UpdateProduct(util.String2Uint(util.GetPathVariable(r, "id")), mapper.ProductDTOToProduct(request, util.GetLoggedUserIDFromToken(r)))
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"UpdateProduct", loggingService)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) GetActiveProducts(w http.ResponseWriter, _ *http.Request) {
	result, err := handler.WSService.GetActiveProducts()
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetActiveProducts", loggingService)
		return
	}
	util.MarshalResult(w, result)
	return
}
