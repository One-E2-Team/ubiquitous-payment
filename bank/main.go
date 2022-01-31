package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
	"ubiquitous-payment/bank/handler"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/bank/repository"
	"ubiquitous-payment/bank/service"
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
		schemaName := "bank_" + os.Getenv("PAN_PREFIX")
		db, err = gorm.Open(mysql.Open(rdb.Username + ":" + rdb.Password + "@tcp(" + rdb.Host + ":" + rdb.Port + ")/" +
			schemaName + "?charset=utf8mb4&parseTime=True&loc=Local"))

		if err != nil {
			fmt.Println("Cannot connect to database! Sleeping 10s and then retrying....")
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("Connected to the database.")
			break
		}
	}

	err = db.AutoMigrate(&model.Privilege{},
		&model.Role{},
		&model.CreditCard{},
		&model.ClientAccount{},
		&model.Client{},
		&model.Transaction{},
		&model.PccOrder{})
	if err != nil {
		return nil
	}

	return db
}

func initRepo(database *gorm.DB) *repository.Repository {
	return &repository.Repository{Database: database}
}

func initService(repo *repository.Repository) *service.Service {
	return &service.Service{BankRepository: repo}
}

func initHandler(service *service.Service) *handler.Handler {
	return &handler.Handler{BankService: service}
}

func handleFunc(handler *handler.Handler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", handler.Test).Methods(http.MethodGet)
	router.HandleFunc("/psp-request", handler.PspRequest).Methods(http.MethodPost)
	fmt.Println("Starting server..")
	host, port := util.GetBankHostAndPort()
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
	bankService := initService(repo)
	bankHandler := initHandler(bankService)
	handleFunc(bankHandler)
}