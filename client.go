package client

import (
	"fmt"
	"net"
)

func Connect(host string, port int) (net.Conn, error) {
	redisHost := fmt.Sprintf("%s:%v", host, port)

	conn, err := net.Dial("tcp", redisHost)
	if err != nil {
		return nil, fmt.Errorf("error connecting to %s: %v", redisHost, err)
	}

	conn.Close()

	return conn, err
}