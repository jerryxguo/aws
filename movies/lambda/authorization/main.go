package main

import (
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	jwt "github.com/dgrijalva/jwt-go"
)

func authorizedToken(secret, tokenString string) bool {
	tokenWithClaim, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err == nil && tokenWithClaim.Valid {
		if _, ok := tokenWithClaim.Claims.(*jwt.StandardClaims); ok {
			return true
		}
	}
	return false
}

func generatePolicy(principalID, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}
	return authResponse
}

func handleRequest(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	secretKey := os.Getenv("API_KEY")
	token := request.AuthorizationToken
	tokenSlice := strings.Split(token, " ")
	var bearerToken string
	if len(tokenSlice) > 1 {
		bearerToken = tokenSlice[len(tokenSlice)-1]

		if allow := authorizedToken(secretKey, bearerToken); allow {
			return generatePolicy("user", "Allow", request.MethodArn), nil
		} else {
			return generatePolicy("user", "Deny", request.MethodArn), nil
		}
	} else {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response
	}
}

func main() {
	lambda.Start(handleRequest)
}
