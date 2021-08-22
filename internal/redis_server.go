package internal

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartRedisServer(ctx context.Context) (testcontainers.Container, string, int, error) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", 0, err
	}

	host, err := redisC.Host(ctx)
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get the redis host: %v", err)
	}

	port, err := redisC.MappedPort(ctx, "6379")
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get the redis port: %v", err)
	}

	return redisC, host, port.Int(), nil
}
