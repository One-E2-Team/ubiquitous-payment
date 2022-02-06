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
	if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
		return "https"
	}
	return "http"
}

func GetInternalPSPProtocol() string {
	return "http"
}

func GetWebShopProtocol() string {
	if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
		return "https"
	}
	return "http"
}

func GetBankProtocol() string {
	if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
		return "https"
	}
	return "http"
}

func GetPccProtocol() string {
	if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
		return "https"
	}
	return "http"
}

func GetInternalPSPHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8002"
	if DockerChecker() {
		pspHost = os.Getenv("PSP_HOST")
		pspPort = os.Getenv("PSP_PORT")
	}
	return pspHost, pspPort
}

func GetExternalPSPHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8002"
	if DockerChecker() {
		pspPort = "1081"
		if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
			pspHost = "host1"
		} else {
			pspHost = "host.docker.internal"
		}
	}
	return pspHost, pspPort
}

func GetPSPFrontHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "3001"
	if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
		pspHost = "host1"
	}
	return pspHost, pspPort
}

func GetInternalWebShopHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8001"
	if DockerChecker() {
		pspHost = "webshop"
		pspPort = "8080"
	}
	return pspHost, pspPort
}

func GetWebShopFrontHostAndPort() (string, string) {
	var webshopHost, webshopPort = "localhost", "3000"
	if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
		webshopHost = "host2"
	}
	return webshopHost, webshopPort
}

func GetInternalPccHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8003"
	if DockerChecker() {
		pspHost = "pcc"
		pspPort = "8080"
	}
	return pspHost, pspPort
}

func GetExternalPccHostAndPort() (string, string) {
	var pccHost, pccPort = "localhost", "8003"
	if DockerChecker() {
		pccPort = "10000"
		if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
			pccHost = "host1"
		} else {
			pccHost = "host.docker.internal"
		}
	}
	return pccHost, pccPort
}

func GetInternalBankHostAndPort() (string, string) {
	bankHost, bankPort := "localhost", os.Getenv("BANK_PORT") //TODO: parametrize for docker
	if DockerChecker() {
		bankHost = os.Getenv("BANK_HOST")
		bankPort = "8080"
	}
	return bankHost, bankPort
}

func GetExternalBankHostAndPort() (string, string) {
	bankHost, bankPort := "localhost", os.Getenv("BANK_PORT")
	if DockerChecker() {
		bankPort = os.Getenv("BANK_EXTERNAL_PORT")
		if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
			bankHost = "host1"
		} else {
			bankHost = "host.docker.internal"
		}
	}
	return bankHost, bankPort
}

func GetBankFrontHostAndPort() (string, string) {
	port := "3002"
	if _, ok := os.LookupEnv("DOCKER_ENV_SET_PROD"); ok {
		return "host1", port
	}
	return "localhost", port
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
		return DatabaseData{Host: "rdb", Port: os.Getenv("RDB_PORT"),
			Username: os.Getenv("RDB_USERNAME"), Password: os.Getenv("RDB_PASSWORD")}
	}
	return DatabaseData{Host: "localhost", Port: rdbPort, Username: "root", Password: "root"}
}
