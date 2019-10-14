package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

var (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI0MzQ2NzAzOTgsImp0aSI6InVzZXIiLCJpc3MiOiJpc3N1ZXIifQ.MSys5QXzjp4_jBEMXC9uxfXAXg2qxg120Q2Zvs4BDCU"
)

func init() {
	err := os.Setenv("API_KEY", "f9c611d65288f1a7135f1c4c1e56fb62")
	if err != nil {
		fmt.Println(err)
	}
	key := os.Getenv("API_KEY")
	fmt.Println(key)
}
func TestHandler(t *testing.T) {
	resp, err := handleRequest(events.APIGatewayCustomAuthorizerRequest{
		AuthorizationToken: "bearer " + token,
		MethodArn:          "arn:",
	})
	assert.IsType(t, nil, err)
	assert.Equal(t, "Allow", resp.PolicyDocument.Statement[0].Effect)

	resp, err = handleRequest(events.APIGatewayCustomAuthorizerRequest{
		AuthorizationToken: "bearer " + "Allow",
		MethodArn:          "arn:",
	})
	assert.IsType(t, nil, err)
	assert.Equal(t, "Deny", resp.PolicyDocument.Statement[0].Effect)
}
