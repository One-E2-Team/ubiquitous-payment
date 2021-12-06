package util

import "os"

func DockerChecker() bool {
	_, ok := os.LookupEnv("DOCKER_ENV_SET_PROD") // dev production environment
	_, ok1 := os.LookupEnv("DOCKER_ENV_SET_DEV") // dev front environment
	return ok || ok1
}

func GetPSPHostAndPort() (string, string) {
	var pspHost, pspPort = "localhost", "8001"
	if DockerChecker() {
		pspHost = "psp"
		pspPort = "8080"
	}
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

func NosqlDockerVars() (string, string, string, string) {
	dbHost := "mongo"
	dbPort := "27017"
	dbUsername := os.Getenv("NOSQLDB_USERNAME")
	dbPassword := os.Getenv("NOSQLDB_PASSWORD")
	return dbHost, dbPort, dbUsername, dbPassword
}

func RDBDockerVars() (string, string, string, string) {
	dbHost := "rdb"
	dbPort := "3306"
	dbUsername := os.Getenv("RDB_USERNAME")
	dbPassword := os.Getenv("RDB_PASSWORD")
	return dbHost, dbPort, dbUsername, dbPassword
}
