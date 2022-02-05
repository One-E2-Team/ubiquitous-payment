package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
	"ubiquitous-payment/psp/handler"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/psp/psputil/rbac"
	"ubiquitous-payment/psp/repository"
	"ubiquitous-payment/psp/service"
	"ubiquitous-payment/util"
)

func initDB() *mongo.Client {
	noSQL := util.GetNoSQLData()
	clientOptions := options.Client().ApplyURI("mongodb://" + noSQL.Username + ":" + noSQL.Password + "@" + noSQL.Host + ":" + noSQL.Port)
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
	createCollection(client, psputil.PspDbName, psputil.WebShopCollectionName)
	createCollection(client, psputil.PspDbName, psputil.TransactionsCollectionName)
	createCollection(client, psputil.PspDbName, psputil.PaymentTypesCollectionName)
	createCollection(client, psputil.PspDbName, psputil.AccountsCollectionName)
	createCollection(client, psputil.PspDbName, psputil.UsersCollectionName)
	createCollection(client, psputil.PspDbName, psputil.BanksCollectionName)
}

func createCollection(client *mongo.Client, dbName string, collectionName string) {
	if err := client.Database(dbName).CreateCollection(psputil.EmptyContext, collectionName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Create " + collectionName + " collection success")
	}
}

func initRepo(client *mongo.Client) *repository.Repository {
	repo := &repository.Repository{Client: client}
	err := repo.AddDBConstraints()
	if err != nil {
		fmt.Println("error in adding DB constraints")
		return nil
	}
	return repo
}

func initPlugin(repo *repository.Repository) {
	psputil.PluginInterfaceContext = repo
}

func initService(pspRepo *repository.Repository) *service.Service {
	return &service.Service{PSPRepository: pspRepo}
}

func initHandler(pspService *service.Service) *handler.Handler {
	return &handler.Handler{PSPService: pspService}
}

func handleFunc(handler *handler.Handler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", handler.Test).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/order-id", rbac.PspRbac(handler.GetNewOrderId, psputil.WebShopOrderPrivilegeName)).Methods(http.MethodGet)
	router.HandleFunc("/api/order", rbac.PspRbac(handler.FillTransaction, psputil.WebShopOrderPrivilegeName)).Methods(http.MethodPost)
	router.HandleFunc("/api/psp/payments/{transactionID}", handler.GetAvailablePaymentTypeNames).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/select-payment", handler.SelectPaymentType).Methods(http.MethodPost)
	router.HandleFunc("/api/psp/payment-success", handler.UpdateTransactionSuccess).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/payment-fail", handler.UpdateTransactionFail).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/payment-error", handler.UpdateTransactionError).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/check-for-payment/bitcoin/{transactionID}", handler.CheckForPaymentBitcoin).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/register-web-shop", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/psp/accept/{webShopID}", handler.AcceptWebShop).Methods(http.MethodPatch)   //TODO: add RBAC for admin
	router.HandleFunc("/api/psp/decline/{webShopID}", handler.DeclineWebShop).Methods(http.MethodPatch) //TODO: add RBAC for admin
	router.HandleFunc("/api/psp/login", handler.LogIn).Methods(http.MethodPost)
	router.HandleFunc("/api/psp/access-uuid", rbac.PspRbac(handler.GetAccessUuidForWebShop, psputil.WebShopTokenPrivilegeName)).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/web-shop-login", handler.LoginWebShop).Methods(http.MethodPost)
	router.HandleFunc("/api/transaction/qrcode/{id}", handler.GetDataForQrCode).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/payment-types/my", handler.GetMyPaymentTypes).Methods(http.MethodGet)
	router.HandleFunc("/api/psp/payment-types", handler.UpdateMyPaymentTypes).Methods(http.MethodPut)

	fmt.Println("Starting server..")
	host, port := util.GetInternalPSPHostAndPort()
	deploymentHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{util.Authorization, util.ContentType, "Accept"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete}))(router)
	var err error
	if util.GetInternalPSPProtocol() == "https" {
		err = http.ListenAndServeTLS(host+":"+port, "./conf/certs/pem/"+host+".cert.pem", "./conf/certs/key/"+host+".key.pem", deploymentHandler)
	} else {
		err = http.ListenAndServe(host+":"+port, deploymentHandler)
	}
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
		fmt.Println(p.Name())
	}
}

func main() {
	client := initDB()
	defer closeConnection(client)
	pspRepo := initRepo(client)
	initPlugin(pspRepo)
	testPlugin("paypal")
	pspService := initService(pspRepo)
	rbac.InitRbacService(pspService)
	pspHandler := initHandler(pspService)
	handleFunc(pspHandler)
}
