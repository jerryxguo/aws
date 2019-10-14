package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	jwt "github.com/dgrijalva/jwt-go"
)

type Request struct {
	Secret          string `json:"secret"`
	User            string `json:"user"`
	Issuer          string `json:"issuer"`
	ExpireInSeconds int    `json:"expireInSeconds"`
}

//generateToken to generate a token for the user based on config
func generateToken(userName, secret, issuer string, expired int) (string, error) {
	expire := time.Now().Add(time.Second * time.Duration(expired)).Unix()
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expire,
		Issuer:    issuer,
		Id:        userName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func handleRequest(request Request) (string, error) {
	return generateToken(request.User, request.Secret, request.Issuer, request.ExpireInSeconds)
}

func main() {
	lambda.Start(handleRequest)
}
