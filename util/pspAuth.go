package util

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

var pspJwt string

const PSPExpiresIn = 86400000

var PSPTokenSecret = ""

type PSPTokenClaims struct {
	WebShop string `json:"webShop"`
	jwt.StandardClaims
}

func initPSPToken() {
	env := os.Getenv("WEB_SHOP_JWT_TOKEN_SECRET")
	if env == "" {
		PSPTokenSecret = "token_secret"
	} else {
		PSPTokenSecret = env
	}
}

func SetupPSPAuth(webShop string) error {
	if PSPTokenSecret == "" {
		initPSPToken()
	}
	claims := PSPTokenClaims{WebShop: webShop, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + PSPExpiresIn,
		IssuedAt:  time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	pspJwt, err = token.SignedString([]byte(PSPTokenSecret))
	return err
}

func ValidatePSPToken(r *http.Request, webShops []string) bool {
	if PSPTokenSecret == "" {
		initPSPToken()
	}
	tokenString, err := getToken(r.Header)
	if err != nil {
		fmt.Println(err)
		return false
	}
	token, err := jwt.ParseWithClaims(tokenString, &PSPTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(PSPTokenSecret), nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	if claims, ok := token.Claims.(*PSPTokenClaims); ok && token.Valid {
		for _, webShop := range webShops {
			if claims.WebShop == webShop {
				return true
			}
		}
		return false
	}
	fmt.Println(err)
	return false
}

func PSPRequest(method string, path string, data []byte, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+pspJwt)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}

func PSPAuth(handler func(http.ResponseWriter, *http.Request), webShops []string) func(http.ResponseWriter, *http.Request) {

	finalHandler := func(pass bool) func(http.ResponseWriter, *http.Request) {
		if pass {
			return handler
		} else {
			return func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Header().Set("Content-Type", "application/json")
				_, _ = writer.Write([]byte("{\"status\":\"fail\", \"reason\":\"unauthorized\"}"))
			}
		}
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		if check := ValidatePSPToken(request, webShops); check {
			finalHandler(true)(writer, request)
		} else {
			finalHandler(false)(writer, request)
		}
	}
}
