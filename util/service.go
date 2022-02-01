package util

import "os"

type DatabaseData struct {
	Host     string
	Port     string
	Username string
	Password string
}

func DockerChecker() bool {
	_, ok := os.LookupEnv("DOCKER_ENV_SET_PROD") // dev production environment
	_, ok1 := os.LookupEnv("DOCKER_ENV_SET_DEV") // dev front environment
	return ok || ok1
}

func GetPSPProtocol() string {
	return "http"
}

func GetWebShopProtocol() string {
	return "http"
}

func GetBankProtocol() string {
	return "http"
}

func GetPccProtocol() string {
	return "http"
}

func GetPSPHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8002"
	if DockerChecker() {
		pspHost = "psp"
		pspPort = "8080"
	}
	return pspHost, pspPort
}

func GetPSPFrontHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "3001"
	return pspHost, pspPort
}

func GetWebShopHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8001"
	if DockerChecker() {
		pspHost = "webshop"
		pspPort = "8080"
	}
	return pspHost, pspPort
}

func GetWebShopFrontHostAndPort() (string, string) {
	return "localhost", "3000"
}

func GetPccHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8003"
	if DockerChecker() {
		pspHost = "pcc"
		pspPort = "8080"
	}
	return pspHost, pspPort
}

func GetBankHostAndPort() (string, string) {
	bankHost, bankPort := "localhost", os.Getenv("BANK_PORT") //TODO: parametrize for docker
	if DockerChecker() {
		bankHost = os.Getenv("BANK_HOST")
		bankPort = "8080"
	}
	return bankHost, bankPort
}

func GetNoSQLData() DatabaseData {
	noSQLPort := "27017"
	if DockerChecker() {
		return DatabaseData{Host: "mongo", Port: noSQLPort,
			Username: os.Getenv("NOSQLDB_USERNAME"), Password: os.Getenv("NOSQLDB_PASSWORD")}
	}
	return DatabaseData{Host: "localhost", Port: noSQLPort, Username: "root", Password: "root"}
}

func GetRDBData() DatabaseData {
	rdbPort := "3306"
	if DockerChecker() {
		return DatabaseData{Host: "rdb", Port: rdbPort,
			Username: os.Getenv("RDB_USERNAME"), Password: os.Getenv("RDB_PASSWORD")}
	}
	return DatabaseData{Host: "localhost", Port: rdbPort, Username: "root", Password: "root"}
}
