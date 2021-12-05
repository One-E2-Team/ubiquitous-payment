package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/handler/mapper"
)

func (handler *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loggedUserId := util.GetLoggedUserIDFromToken(r)
	//if loggedUserId == 0 {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	err = handler.WSService.CreateProduct(mapper.ProductDTOToProduct(request, loggedUserId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (handler *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pathVars := mux.Vars(r)
	err = handler.WSService.UpdateProduct(util.String2Uint(pathVars["id"]), mapper.ProductDTOToProduct(request, 0))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
