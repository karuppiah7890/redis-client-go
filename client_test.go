package client_test

import (
	"testing"

	client "github.com/karuppiah7890/redis-client-go"
)

func TestConnect(t *testing.T) {
	err := client.Connect("localhost", 6379)
	if err != nil {
		t.Errorf("Connection to Redis Server failed: %v", err)
	}
}
