package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/handler/mapper"
)

func (handler *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	loggedUserId := util.GetLoggedUserIDFromToken(r)
	if loggedUserId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := r.ParseMultipartForm(0); err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	var request dto.ProductDTO
	data := r.MultipartForm.Value["data"]
	err := json.Unmarshal([]byte(data[0]), &request)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	err = handler.WSService.CreateProduct(mapper.ProductDTOToProduct(request, loggedUserId), r)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (handler *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}

	pathVars := mux.Vars(r)
	err = handler.WSService.UpdateProduct(util.String2Uint(pathVars["id"]), mapper.ProductDTOToProduct(request, 0))
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) GetActiveProducts(w http.ResponseWriter, r *http.Request) {
	result, err := handler.WSService.GetActiveProducts()
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}
	util.MarshalResult(w, result)
	return
}