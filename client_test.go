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

	conn, err := client.Connect(host, port)
	if err != nil {
		t.Errorf("Connection to Redis Server failed: %v", err)
	}

	conn.Close()
	numberOfConnectionsReceived := server.NumberOfConnectionsReceived()

	if numberOfConnectionsReceived != 1 {
		t.Errorf("Expected 1 connection to be received by the Redis Server but got: %v", numberOfConnectionsReceived)
	}
}

func TestPing(t *testing.T) {
	ctx := context.Background()
	redisC, err := internal.StartRedisServer(ctx)
	if err != nil {
		t.Errorf("failed to start the redis container: %v", err)
		return
	}

	defer redisC.Terminate(ctx)

	// Maybe move this logic inside StartRedisServer
	host, err := redisC.Host(ctx)
	if err != nil {
		t.Errorf("failed to get the redis host: %v", err)
		return
	}

	// Maybe move this logic inside StartRedisServer
	port, err := redisC.MappedPort(ctx, "6379")
	if err != nil {
		t.Errorf("failed to get the redis port: %v", err)
		return
	}

	conn, err := client.Connect(host, port.Int())
	if err != nil {
		t.Errorf("Connection to Redis Server failed: %v", err)
		return
	}

	conn.Close()
}
