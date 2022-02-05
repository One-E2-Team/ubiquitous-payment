package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/handler"
	"ubiquitous-payment/webshop/model"
	"ubiquitous-payment/webshop/repository"
	"ubiquitous-payment/webshop/service"
	"ubiquitous-payment/webshop/wsutil/pspAuth"
	"ubiquitous-payment/webshop/wsutil/rbac"
)

func initDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	time.Sleep(5 * time.Second)
	rdb := util.GetRDBData()
	for {
		db, err = gorm.Open(mysql.Open(rdb.Username + ":" + rdb.Password + "@tcp(" + rdb.Host + ":" + rdb.Port + ")/webshop?charset=utf8mb4&parseTime=True&loc=Local"))

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
		&model.User{},
		&model.Profile{},
		&model.Account{},
		&model.Order{},
		&model.PaymentType{},
		&model.Product{},
		&model.PSPOrder{},
		&model.WebShop{})
	if err != nil {
		return nil
	}

	return db
}

func initRepo(database *gorm.DB) *repository.Repository {
	return &repository.Repository{RelationalDatabase: database}
}

func initService(wsRepo *repository.Repository) *service.Service {
	return &service.Service{WSRepository: wsRepo}
}

func initHandler(wsService *service.Service) *handler.Handler {
	return &handler.Handler{WSService: wsService}
}

func handleFunc(handler *handler.Handler) {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../staticdata/"))))
	router.HandleFunc("/test", handler.Test).Methods(http.MethodGet)
	router.HandleFunc("/api/login", handler.LogIn).Methods(http.MethodPost)
	router.HandleFunc("/api/users", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/products", handler.GetActiveProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products",
		rbac.WebShopRbac(handler.CreateProduct, "CREATE_PRODUCT")).Methods(http.MethodPost)
	router.HandleFunc("/api/products/{id}",
		rbac.WebShopRbac(handler.UpdateProduct, "UPDATE_PRODUCT")).Methods(http.MethodPut)
	router.HandleFunc("/api/orders/{id}",
		rbac.WebShopRbac(handler.CreateOrder, "CREATE_ORDER")).Methods(http.MethodPost)
	router.HandleFunc("/api/psp-access-token", handler.SetPSPAccessToken).Methods(http.MethodPost)
	router.HandleFunc("/api/payment-types", handler.GetValidPaymentTypes).Methods(http.MethodGet)
	router.HandleFunc("/api/accounts/{name}", handler.GetAccountsByPaymentType).Methods(http.MethodGet)
	router.HandleFunc("/api/accounts/{id}", handler.UpdateAccount).Methods(http.MethodPut)
	router.HandleFunc("/api/accounts", handler.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/accounts/{id}", handler.DeleteAccount).Methods(http.MethodDelete)
	router.HandleFunc("/api/my-orders", handler.GetMyOrders).Methods(http.MethodGet)
	router.HandleFunc("/api/psp-order/{id}/{status}", handler.UpdatePspOrder).Methods(http.MethodPut)
	router.HandleFunc("/api/confirm-password", handler.ConfirmPassword).Methods(http.MethodPut)
	router.HandleFunc("/api/my-orders/seller", handler.GetSellersOrders).Methods(http.MethodGet)
	fmt.Println("Starting server..")
	host, port := util.GetWebShopHostAndPort()
	var err error
	err = http.ListenAndServe(host+":"+port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{util.Authorization, util.ContentType, "Accept"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete}))(router))
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
	db := initDB()
	wsRepo := initRepo(db)
	wsService := initService(wsRepo)
	accessToken, err := wsService.GetPSPAccessToken()
	if err != nil {
		accessToken = ""
	}
	pspAuth.SetPspAccessToken(accessToken)
	rbac.InitRbacService(wsService)
	wsHandler := initHandler(wsService)
	handleFunc(wsHandler)
}
