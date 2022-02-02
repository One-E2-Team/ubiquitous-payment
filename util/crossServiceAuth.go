package util

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

var csJwt string

const CsTokenExpiresIn = 86400000

var CsTokenSecret = ""

type CsTokenClaims struct {
	Service string `json:"service"`
	jwt.StandardClaims
}

func initCsTokenSecret() {
	env := os.Getenv("CROSS_SERVICE_JWT_TOKEN_SECRET")
	if env == "" {
		CsTokenSecret = "token_secret"
	} else {
		CsTokenSecret = env
	}
}

func SetupCsAuth(serviceName string) error {
	if CsTokenSecret == "" {
		initCsTokenSecret()
	}
	claims := CsTokenClaims{Service: serviceName, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + CsTokenExpiresIn,
		IssuedAt:  time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	csJwt, err = token.SignedString([]byte(CsTokenSecret))
	return err
}

func ValidateCsToken(r *http.Request, services []string) (bool, string) {
	if CsTokenSecret == "" {
		initCsTokenSecret()
	}
	tokenString, err := getToken(r.Header)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	token, err := jwt.ParseWithClaims(tokenString, &CsTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(CsTokenSecret), nil
	})
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	if claims, ok := token.Claims.(*CsTokenClaims); ok && token.Valid {
		for _, value := range services {
			if claims.Service == value {
				return true, claims.Service
			}
		}
		return false, ""
	} else {
		fmt.Println(err)
		return false, ""
	}
}

func CrossServiceRequest(method string, path string, data []byte, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set(Authorization, "Bearer "+csJwt)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}
