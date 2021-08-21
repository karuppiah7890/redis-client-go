package client_test

import (
	"math/rand"
	"testing"

	client "github.com/karuppiah7890/redis-client-go"
	"github.com/karuppiah7890/redis-client-go/internal"
)

func TestConnect(t *testing.T) {
	host := "localhost"
	port := rand.Intn(65536)
	server := internal.NewMockRedisServer(host, port)
	err := server.Start()
	if err != nil {
		t.Errorf("Starting mock Redis Server failed: %v", err)
	}
	defer server.Stop()

	_, err = client.Connect(host, port)
	if err != nil {
		t.Errorf("Connection to Redis Server failed: %v", err)
	}

	numberOfConnectionsReceived := server.NumberOfConnectionsReceived()

	if numberOfConnectionsReceived != 1 {
		t.Errorf("Expected 1 connection to be received by the Redis Server but got: %v", numberOfConnectionsReceived)
	}
}
