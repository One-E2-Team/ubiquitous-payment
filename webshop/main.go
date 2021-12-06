package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/handler"
	"ubiquitous-payment/webshop/model"
	"ubiquitous-payment/webshop/repository"
	"ubiquitous-payment/webshop/service"
	"ubiquitous-payment/webshop/wsutil"
)

func initDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	time.Sleep(5 * time.Second)
	var dbHost, dbPort, dbUsername, dbPassword = "localhost", "3306", "root", "root"
	_, ok := os.LookupEnv("DOCKER_ENV_SET_PROD")
	if ok {
		dbHost = "rdb"
		dbPort = "3306"
		dbUsername = os.Getenv("RDB_USERNAME")
		dbPassword = os.Getenv("RDB_PASSWORD")
	}
	for {
		db, err = gorm.Open(mysql.Open(dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/webshop?charset=utf8mb4&parseTime=True&loc=Local"))

		if err != nil {
			fmt.Println("Cannot connect to database! Sleeping 10s and then retrying....")
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("Connected to the database.")
			break
		}
	}

	err = db.AutoMigrate(&model.Privilege{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.Role{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.Profile{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.Account{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.Order{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.PaymentType{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.PSPOrder{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&model.WebShop{})
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
	router.HandleFunc("/test", handler.Test).Methods(util.HttpGet)
	router.HandleFunc("/api/login", handler.LogIn).Methods(util.HttpPost)
	router.HandleFunc("/api/users", handler.Register).Methods(util.HttpPost)
	router.HandleFunc("/api/products",
		wsutil.RBAC(handler.CreateProduct, "CREATE_PRODUCT", false)).Methods(util.HttpPost)
	router.HandleFunc("/api/products/{id}",
		wsutil.RBAC(handler.UpdateProduct, "UPDATE_PRODUCT", false)).Methods(util.HttpPut)
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
	db := initDB()
	wsRepo := initRepo(db)
	wsService := initService(wsRepo)
	wsutil.InitRbacService(wsService)
	wsHandler := initHandler(wsService)
	handleFunc(wsHandler)
}
