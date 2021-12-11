package psputil

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const ExpiresIn = 86400

var PSPTokenSecret = ""

type PSPTokenClaims struct {
	LoggedUserId string `json:"loggedUserId"`
	jwt.StandardClaims
}

func initTokenSecret() {
	env := os.Getenv("PUBLIC_JWT_TOKEN_SECRET")
	if env == "" {
		PSPTokenSecret = "token_secret"
	} else {
		PSPTokenSecret = env
	}
}

func CreateToken(userId string, issuer string) (string, error) {
	if PSPTokenSecret == "" {
		initTokenSecret()
	}
	claims := PSPTokenClaims{LoggedUserId: userId, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + ExpiresIn,
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuer,
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(PSPTokenSecret))
}

func getToken(header http.Header) (string, error) {
	reqToken := header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("NO_TOKEN")
	}
	return strings.TrimSpace(splitToken[1]), nil
}

func GetLoggedUserIDFromToken(r *http.Request) string {
	if PSPTokenSecret == "" {
		initTokenSecret()
	}
	tokenString, err := getToken(r.Header)
	if err != nil {
		var err1 error
		tokenString, err1 = getTokenFromParams(r.URL.String())
		if err1 != nil {
			fmt.Println(err, err1)
			return ""
		}
	}
	return getLoggedUserIDFromPureToken(tokenString)
}

func getTokenFromParams(s string) (string, error) {
	err := fmt.Errorf("generic error, no token in URL")
	paramsPart := strings.Split(s, "?")
	if len(paramsPart) < 2 {
		return "", err
	}
	tokenTilEnd := strings.Split(paramsPart[1], "token=")
	if len(paramsPart) < 2 {
		return "", err
	}
	tokenEscaped := strings.Split(tokenTilEnd[1], "&")
	token, err := url.QueryUnescape(tokenEscaped[0])
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}

func getLoggedUserIDFromPureToken(tok string) string {
	token, err := jwt.ParseWithClaims(tok, &PSPTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(PSPTokenSecret), nil
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if claims, ok := token.Claims.(*PSPTokenClaims); ok && token.Valid {
		return claims.LoggedUserId
	}
	fmt.Println("error during parsing token in PSP")
	return ""
}
