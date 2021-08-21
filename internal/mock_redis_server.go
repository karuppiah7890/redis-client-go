package internal

import (
	"fmt"
	"net"
)

type MockRedisServer struct {
	host                 string
	port                 int
	server_type          string
	listener             net.Listener
	connections_received int
	running              bool
}

func NewMockRedisServer(host string, port int) *MockRedisServer {
	return &MockRedisServer{
		host:        host,
		port:        port,
		server_type: "tcp",
		running:     false,
	}
}

func (server *MockRedisServer) Start() error {
	// Listen for incoming connections.
	l, err := net.Listen(server.server_type, fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		return fmt.Errorf("error listening for connections: %v", err.Error())
	}
	server.listener = l
	fmt.Printf("Listening at %s:%d\n", server.host, server.port)
	server.running = true
	go func() {
		for server.running {
			// Listen for an incoming connection.
			conn, err := l.Accept()
			if err != nil {
				fmt.Printf("error accepting connections: %v\n", err.Error())
				continue
			}
			server.connections_received++
			// Handle connections in a new goroutine.
			go handleRequest(conn)
		}
	}()

	return nil
}

func (server *MockRedisServer) NumberOfConnectionsReceived() int {
	return server.connections_received
}

func (server *MockRedisServer) Stop() error {
	server.running = false
	// Close the listener
	err := server.listener.Close()
	if err != nil {
		return fmt.Errorf("error closing listener: %v", err.Error())
	}
	return nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	conn.Close()
}
