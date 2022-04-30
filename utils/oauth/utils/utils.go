package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type SignedDetails struct {
	UserID int64
	jwt.RegisteredClaims
}

var secretKey = []byte("NOMORESAPA")

func ValidateToken(request *http.Request) (claims *SignedDetails, msg string) {
	fmt.Println("Inside here")
	extractedToken, extractErr := ExtractToken(request)
	if extractErr != nil {
		fmt.Println("Insider extrace err")
		fmt.Println(extractErr)
		return nil, ""
	}
	SignedToken := extractedToken
	token, err := jwt.ParseWithClaims(
		SignedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "Invalid Token"
		msg = err.Error()
		return
	}
	if claimErr := claims.Valid(); claimErr != nil {
		msg = "token is invalid"
		msg = err.Error()
		return
	}
	return claims, msg
}
func ExtractToken(request *http.Request) (string, error) {
	bearerToken := request.Header.Get("Authorization")
	tokenStr := strings.Split(bearerToken, " ")
	if len(tokenStr) != 2 {
		return "", errors.New("invalid token")
	}
	return tokenStr[1], nil
}
