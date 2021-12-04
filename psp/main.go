package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ubiquitous-payment/psp/handler"
	"ubiquitous-payment/psp/repository"
	"ubiquitous-payment/psp/service"
)

func initRepo() *repository.Repository {
	return &repository.Repository{}
}

func initService(pspRepo *repository.Repository) *service.Service {
	return &service.Service{PSPRepository: pspRepo}
}

func initHandler(pspService *service.Service) *handler.Handler {
	return &handler.Handler{PSPService: pspService}
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
