package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/webshop/handler"
	"ubiquitous-payment/webshop/repository"
	"ubiquitous-payment/webshop/service"
)

func initRepo() *repository.Repository {
	return &repository.Repository{}
}

func initService(wsRepo *repository.Repository) *service.Service {
	return &service.Service{WSRepository: wsRepo}
}

func initHandler(wsService *service.Service) *handler.Handler {
	return &handler.Handler{WSService: wsService}
}

func handleFunc(handler *handler.Handler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", handler.Test).Methods("GET")
	fmt.Println("Starting server..")
	host := "localhost"
	port := "81"
	var err error
	err = http.ListenAndServe(host+":"+port, router)
	/*host, port := util.GetConnectionHostAndPort()
	if util.DockerChecker() {
		err = http.ListenAndServeTLS(":"+port, "../cert.pem", "../key.pem", router)
	} else {
		err = http.ListenAndServe(host+":"+port, router)
	}*/
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	pspRepo := initRepo()
	pspService := initService(pspRepo)
	pspHandler := initHandler(pspService)
	handleFunc(pspHandler)
}