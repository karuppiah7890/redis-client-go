package client

import (
	"bytes"
	"fmt"
	"net"
)

func Connect(host string, port int) (net.Conn, error) {
	redisHost := fmt.Sprintf("%s:%v", host, port)

	conn, err := net.Dial("tcp", redisHost)
	if err != nil {
		return nil, fmt.Errorf("error connecting to %s: %v", redisHost, err)
	}

	return conn, err
}

func Ping(conn net.Conn) (string, error) {
	ping := "PING\r\n"
	n, err := conn.Write([]byte(ping))

	if err != nil {
		return "", fmt.Errorf("error while pinging: %v", err)
	}

	if n != len(ping) {
		return "", fmt.Errorf("error while pinging. not all bytes were written to connection. expected to write: %v bytes, but wrote: %v bytes", len(ping), n)
	}

	buf := make([]byte, 512)

	_, err = conn.Read(buf)

	if err != nil {
		return "", fmt.Errorf("error while pinging: %v", err)
	}

	if buf[0] != '+' {
		return "", fmt.Errorf("error while pinging. expected simple string but got something else. first byte: %v", buf[0])
	}

	if !bytes.Equal(buf[1:5], []byte("PONG")) {
		return "", fmt.Errorf("error while pinging. expected pong as response but got something else. response: %v", string(buf))
	}

	return "PONG", nil
}
