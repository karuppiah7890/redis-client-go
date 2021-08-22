package client_test

import (
	"context"
	"math/rand"
	"testing"

	client "github.com/karuppiah7890/redis-client-go"
	"github.com/karuppiah7890/redis-client-go/internal"
)

func TestConnect(t *testing.T) {
	t.Skip()
	host := "localhost"
	port := rand.Intn(65536)
	server := internal.NewMockRedisServer(host, port)
	err := server.Start()
	if err != nil {
		t.Errorf("Starting mock Redis Server failed: %v", err)
	}
	defer server.Stop()

	redisClient, err := client.Connect(host, port)
	if err != nil {
		t.Errorf("Connection to Redis Server failed: %v", err)
	}

	redisClient.Close()
	numberOfConnectionsReceived := server.NumberOfConnectionsReceived()

	if numberOfConnectionsReceived != 1 {
		t.Errorf("Expected 1 connection to be received by the Redis Server but got: %v", numberOfConnectionsReceived)
	}
}

func TestClient(t *testing.T) {
	ctx := context.Background()
	redisC, host, port, err := internal.StartRedisServer(ctx)
	if err != nil {
		t.Errorf("failed to start the redis container: %v", err)
		return
	}

	defer redisC.Terminate(ctx)

	redisClient, err := client.Connect(host, port)
	if err != nil {
		t.Errorf("Connection to Redis Server failed: %v", err)
		return
	}

	defer redisClient.Close()

	t.Run("TestPing", func(t *testing.T) {
		pingResponse, err := redisClient.Ping()
		if err != nil {
			t.Errorf("Expected no errors in PING but got: %v", err)
		}

		if pingResponse != "PONG" {
			t.Errorf("Expected PONG as reply for PING but got: %v", pingResponse)
		}
	})

	t.Run("TestExecuteCommand", func(t *testing.T) {
		t.Run("Running Non existent command", func(t *testing.T) {
			response, err := redisClient.ExecuteCommand("BLAH")
			if err == nil {
				t.Errorf("Expected errors in executing BLAH command but got response: %v", response)
			}

			if response != "" {
				t.Errorf("Expected empty string as reply for BLAH but got: %v", response)
			}
		})
	})
}
