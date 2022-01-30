package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
	"ubiquitous-payment/pcc/handler"
	"ubiquitous-payment/pcc/model"
	"ubiquitous-payment/pcc/repository"
	"ubiquitous-payment/pcc/service"
	"ubiquitous-payment/util"
)

func initDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	time.Sleep(5 * time.Second)
	rdb := util.GetRDBData()
	for {
		db, err = gorm.Open(mysql.Open(rdb.Username + ":" + rdb.Password + "@tcp(" + rdb.Host + ":" + rdb.Port + ")/pcc?charset=utf8mb4&parseTime=True&loc=Local"))

		if err != nil {
			fmt.Println("Cannot connect to database! Sleeping 10s and then retrying....")
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("Connected to the database.")
			break
		}
	}

	err = db.AutoMigrate(&model.Bank{},
		&model.PccOrder{})
	if err != nil {
		return nil
	}

	return db
}

func initRepo(database *gorm.DB) *repository.Repository {
	return &repository.Repository{RelationalDatabase: database}
}

func initService(repo *repository.Repository) *service.Service {
	return &service.Service{Repository: repo}
}

func initHandler(service *service.Service) *handler.Handler {
	return &handler.Handler{Service: service}
}

func handleFunc(handler *handler.Handler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", handler.Test).Methods(http.MethodGet)
	fmt.Println("Starting server..")
	host, port := util.GetPccHostAndPort()
	var err error
	err = http.ListenAndServe(host+":"+port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{util.Authorization, util.ContentType, "Accept"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete}))(router))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	db := initDB()
	repo := initRepo(db)
	service := initService(repo)
	handler := initHandler(service)
	handleFunc(handler)
}