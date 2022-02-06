package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"os"
	"time"
	"ubiquitous-payment/bank/bankutil/rbac"
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
		&model.Test{})
	if err != nil {
		return nil
	}

	return db
}

func initRepo(database *gorm.DB) *repository.Repository {
	return &repository.Repository{Database: database}
}

func initService(repo *repository.Repository) *service.Service {
	return &service.Service{Repository: repo}
}

func initHandler(service *service.Service) *handler.Handler {
	return &handler.Handler{BankService: service}
}

func handleFunc(handler *handler.Handler) {
	router := mux.NewRouter().StrictSlash(true)

	//public API
	router.HandleFunc("/test", handler.Test).Methods(http.MethodGet)
	router.HandleFunc("/test-encryption", handler.TestEncryption).Methods(http.MethodPost)
	router.HandleFunc("/api/clients", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", handler.LogIn).Methods(http.MethodPost)
	router.HandleFunc("/api/pay/{payment-url-id}", handler.Pay).Methods(http.MethodPost)
	router.HandleFunc("/api/payment-details/{payment-url-id}", handler.GetPaymentDetails).Methods(http.MethodGet)

	//psp calls
	router.HandleFunc("/psp-request", handler.PspRequest).Methods(http.MethodPost)
	router.HandleFunc("/api/payment-check/{id}", handler.CheckPayment).Methods(http.MethodGet)

	//pcc calls
	router.HandleFunc("/pcc-issuer-pay",
		rbac.PccRbac(handler.IssuerPay, "pcc")).Methods(http.MethodPost)

	// logged users API
	router.HandleFunc("/api/account",
		rbac.BankRbac(handler.GetMyAccount, "READ_ACCOUNT")).Methods(http.MethodGet)
	router.HandleFunc("/api/confirm-password",
		rbac.BankRbac(handler.ConfirmPassword, "CHECK_PASSWORD")).Methods(http.MethodPost)
	router.HandleFunc("/api/transactions",
		rbac.BankRbac(handler.GetMyTransactions, "READ_TRANSACTIONS")).Methods(http.MethodGet)
	router.HandleFunc("/api/all-transactions",
		rbac.BankRbac(handler.GetAllTransactions, "READ_ALL_TRANSACTIONS")).Methods(http.MethodGet)
	fmt.Println("Starting server..")
	host, port := util.GetInternalBankHostAndPort()
	deploymentHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{util.Authorization, util.ContentType, "Accept"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete}))(router)
	var err error
	if util.GetBankProtocol() == "https" {
		err = http.ListenAndServeTLS(host+":"+port, "./conf/certs/pem/"+host+".cert.pem", "./conf/certs/key/"+host+".key.pem", deploymentHandler)
	} else {
		err = http.ListenAndServe(host+":"+port, deploymentHandler)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}

func checkAccountsActivity(repo *repository.Repository) {
	for {
		allClients, err := repo.GetAllClients()
		if err != nil {
			fmt.Println(err)
		} else {
			for i, client := range allClients {
				if client.LastActivityTimestamp.Before(time.Now().Add(-(time.Hour * 2160))) { //90 days
					allClients[i].IsDeleted = true
					err = repo.Update(&allClients[i])
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
		time.Sleep(time.Hour * 24)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	db := initDB()
	repo := initRepo(db)
	bankService := initService(repo)
	bankHandler := initHandler(bankService)
	_ = util.SetupCsAuth("bank")
	rbac.InitRbacService(bankService)

	go checkAccountsActivity(repo)

	handleFunc(bankHandler)
}
