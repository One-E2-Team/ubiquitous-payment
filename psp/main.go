package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
	"ubiquitous-payment/psp/handler"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/psp/repository"
	"ubiquitous-payment/psp/service"
	"ubiquitous-payment/util"
)

func initDB() *mongo.Client {
	var dbHost, dbPort, dbUsername, dbPassword = "localhost", "27017", "root", "root"
	if util.DockerChecker() {
		dbHost, dbPort, dbUsername, dbPassword = util.NosqlDockerVars()
	}
	clientOptions := options.Client().ApplyURI("mongodb://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort)
	for {
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			fmt.Println("Cannot connect to MongoDB! Sleeping 10s and then retrying....")
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("Connected to MongoDB")
			initCollections(client)
			return client
		}
	}
}

func initCollections(client *mongo.Client) {
	const webshopsCollectionName = "psp-clients"
	const transactionsCollectionName = "psp-transactions"
	const pspDbName = "psp-db"
	createCollection(client, pspDbName, webshopsCollectionName)
	createCollection(client, pspDbName, transactionsCollectionName)
}

func createCollection(client *mongo.Client, dbName string, collectionName string) {
	if err := client.Database(dbName).CreateCollection(context.TODO(), collectionName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Create " + collectionName + " collection success")
	}
}

func initRepo(client *mongo.Client) *repository.Repository {
	return &repository.Repository{Client: client}
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
	host, port := util.GetPSPHostAndPort()
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

func closeConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("Failed to close MongoDB.")
		return
	}
	fmt.Println("Connection to MongoDB closed.")
}

func testPlugin(pluginName string) {
	p, err := psputil.GetPlugin(pluginName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p.Test())
	}
}

func main() {
	testPlugin("paypal")
	client := initDB()
	defer closeConnection(client)
	pspRepo := initRepo(client)
	pspService := initService(pspRepo)
	pspHandler := initHandler(pspService)
	handleFunc(pspHandler)
}
