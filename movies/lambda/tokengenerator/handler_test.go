package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	token, err := handleRequest(Request{
		Secret:          "f9c611d65288f1a7135f1c4c1e56fb62",
		User:            "user",
		Issuer:          "issuer",
		ExpireInSeconds: 3600 * 24 * 10000,
	})
	assert.IsType(t, nil, err)
	assert.True(t, len(token) > 100)
	fmt.Println(token)
}
