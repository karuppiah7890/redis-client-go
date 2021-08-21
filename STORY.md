# Story

So, this is the story of how I'm building a Redis Client library in Golang

https://github.com/karuppiah7890/redis-client-go

It's not gonna be some mainstream library. I'm not aiming for that, though that would be cool and all, that's a lot of work and lot of maintenance. I'm just aiming to try and learn Redis protocol by working on a project like this

This is purely for learning, and fun and to feed my boredom. And I'm not sure if I can cover the complete Redis protocol. So clearly I'm not planning to aim to write a full blown library to maintain and work on in the future. Just fun fun fun for now and learning of course

I was looking at what kind of interface the other libraries provide. I started thinking from the naming perspective and even the versioning. But I figured I could do away with those since I'm not aiming for this project to be used by others. But it was cool to notice the naming and stuff. Like, I was wondering if using `redis-client-go` would be okay and if I would have to ask folks to do something like

```go
import "github.com/karuppiah7890/redis-client-go/client"
```

Or

```go
import "github.com/karuppiah7890/redis-client-go/redisclient"
```

But looks like people have named their top level packages in a nice way and that gets used regardless of the repo name

For example, https://github.com/go-redis/redis uses code like this -

```go
import (
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
...
...
```

Notice how it does a `/v8` I think that's versioning? I can see the v8.x.y in releases too https://github.com/go-redis/redis/releases , https://github.com/go-redis/redis/releases/tag/v8.11.2 , https://github.com/go-redis/redis/releases/tag/v8.11.3

I was also looking at other Redis clients in Go by checking https://redis.io/clients#go

https://github.com/alphazero/Go-Redis

https://github.com/insmo/godis

https://github.com/piaohao/godis

https://github.com/keimoon/gore

https://github.com/xuyu/goredis

https://github.com/go-redis/redis

https://github.com/stfnmllr/go-resp3

I just skimmed through the name :P Didn't go through the whole docs etc. Noticed some `import` and one or two lines of code here and there. Realized there's auth stuff to support too ;) `AUTH` in case the connection URL contains password. And maybe even ACL stuff, in case there's username AND password

I noticed RESP stuff - https://github.com/antirez/RESP3

I remember learning that RESP is the Redis Protocol. Redis Serialization Protocol.

I'll be trying to implement that protocol and trying to see if it works and how it works and stuff

I also had some parallel ideas and also some ideas about this project which is below -

Redis client library in Golang

Redis client library in Rust

Redis CLI client in Golang

Redis CLI client in Rust

All basic stuff

Write simple small tests, with mocks / mock servers. Run those tests. Run in GitHub Actions? ;) Also, use actual Redis server to run Integration tests :)

Have stories for these. We can put library and CLI in same repo for each language. But maintain separation of concerns.

Support for Redis Cluster? Sure

Learn RESP protocol in the journey. And some basic socket programming in Golang and Rust too in the future

Maybe also learn the Redis Cluster Bus protocol, simply. And create a client for it, simply. To try it out! :)

How about creating a Redis Server? Hmm? In Golang and Rust. A mock server? Just basic one. With implementation of basic stuff. Like GET, SET

How about creating CLIs to deal with RDB and AOF? ;) Hmm

---

That's that! I guess I'll get started by reading some RESP stuff! And using `telnet` and what not! ;)

I'll also check about `AUTH` command and auth stuff, and also about "persisting connections" ;)

I'm checking about RESP now

https://duckduckgo.com/?t=ffab&q=resp+protocol&ia=web

https://github.com/antirez/RESP3

https://redis.io/topics/protocol

So, it's a TCP based protocol. Of course. https://redis.io/topics/protocol#networking-layer

I'll have to create a simple TCE connection to the Redis Server first. Hmm

```bash
redis-client-go $ gst
On branch main
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	client.go
	client_test.go

nothing added to commit but untracked files present (use "git add" to track)
redis-client-go $ git remote -v
origin	git@github.com:karuppiah7890/redis-client-go.git (fetch)
origin	git@github.com:karuppiah7890/redis-client-go.git (push)
redis-client-go $ go mod init github.com/karuppiah7890/redis-client-go
go: creating new go.mod: module github.com/karuppiah7890/redis-client-go
go: to add module requirements and sums:
	go mod tidy
redis-client-go $ gst
On branch main
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	client.go
	client_test.go
	go.mod

nothing added to commit but untracked files present (use "git add" to track)
redis-client-go $ go mod tidy
redis-client-go $ gst
On branch main
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	client.go
	client_test.go
	go.mod

nothing added to commit but untracked files present (use "git add" to track)
redis-client-go $ go mod tidy -v
redis-client-go $ gst
On branch main
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	client.go
	client_test.go
	go.mod

nothing added to commit but untracked files present (use "git add" to track)
redis-client-go $ code .
redis-client-go $ 
```

I'm starting to write tests

```bash
redis-client-go $ go test -v
=== RUN   TestConnect
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	0.376s
redis-client-go $ go test -v ./...
=== RUN   TestConnect
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	0.092s
```

It doesn't check anything as of now

I'm just writing a function called as `Connect` for now, to connect to a TCP server running at a host and port

```go
package client

func Connect(host string, port int) {

}
```

```go
package client_test

import (
	"testing"

	client "github.com/karuppiah7890/redis-client-go"
)

func TestConnect(t *testing.T) {
	client.Connect("localhost", 6379)
}
```

---

Now I'm making it a bit better

```go
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
```

```go
package client

func Connect(host string, port int) error {
	return nil
}
```

I also added a `Makefile` ! :)

```Makefile
test:
	go test -v ./...
```

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	0.375s
```

---

Next I need to write a dummy TCP server for the tests

https://duckduckgo.com/?t=ffab&q=golang+tcp+server&ia=web

https://duckduckgo.com/?t=ffab&q=golang+tcp&ia=web

https://duckduckgo.com/?t=ffab&q=golang+net+tcp&ia=web

https://pkg.go.dev/net

https://pkg.go.dev/net#Dial

Lot of details. From what I read, I think I can use `tcp` network for both IPv4 and IPv6. I think IPv4 is good for now as I'm not gonna test all the IP version protocol stuff. I can dial to my local too :) with just ":80" etc I think

https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

```go
package main

import (
    "fmt"
    "net"
    "os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  conn.Close()
}
```

I made some modifications and tried to run the test. Funnily I made the test block, lol

```go
package internal

import (
	"fmt"
	"net"
)

type MockRedisServer struct {
	host        string
	port        int
	server_type string
	listener    net.Listener
}

func NewMockRedisServer(host string, port int) *MockRedisServer {
	return &MockRedisServer{
		host:        host,
		port:        port,
		server_type: "tcp",
	}
}

func (server *MockRedisServer) Start() error {
	// Listen for incoming connections.
	l, err := net.Listen(server.server_type, fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		return fmt.Errorf("error listening for connections: %v", err.Error())
	}
	fmt.Printf("Listening at %s:%d\n", server.host, server.port)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            return fmt.Errorf("error accepting connections: %v", err.Error())
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }

	return nil
}

func (server *MockRedisServer) Stop() error {
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
```

```go
package client_test

import (
	"testing"

	client "github.com/karuppiah7890/redis-client-go"
	"github.com/karuppiah7890/redis-client-go/internal"
)

func TestConnect(t *testing.T) {
	host := "localhost"
	port := 6379
	server := internal.NewMockRedisServer(host, port)
	err := server.Start()
	if err != nil {
		t.Errorf("Starting mock Redis Server failed: %v", err)
	}
	defer server.Stop();

	err = client.Connect(host, port)
	if err != nil {
		t.Errorf("Connection to Redis Server failed: %v", err)
	}
}
```

I avoided the errors from `server.Stop()` actually

```bash
redis-client-go $ make test
go test -v ./...
^C=== RUN   TestConnect
Listening at localhost:6379
FAIL	github.com/karuppiah7890/redis-client-go	5.822s
make: *** [test] Error 1

redis-client-go $ 
```

Funny thing, but I finally changed the blocking call like this -

```go
func (server *MockRedisServer) Start() error {
	// Listen for incoming connections.
	l, err := net.Listen(server.server_type, fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		return fmt.Errorf("error listening for connections: %v", err.Error())
	}
	fmt.Printf("Listening at %s:%d\n", server.host, server.port)
	go func() {
		for {
			// Listen for an incoming connection.
			conn, err := l.Accept()
			if err != nil {
				fmt.Printf("error accepting connections: %v", err.Error())
			}
			// Handle connections in a new goroutine.
			go handleRequest(conn)
		}
	}()

	return nil
}
```

With goroutines

Now, the problem was -

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
--- FAIL: TestConnect (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x28 pc=0x114966e]

goroutine 19 [running]:
testing.tRunner.func1.2(0x11673c0, 0x128ffc0)
	/usr/local/Cellar/go/1.16.6/libexec/src/testing/testing.go:1143 +0x332
testing.tRunner.func1(0xc000186600)
	/usr/local/Cellar/go/1.16.6/libexec/src/testing/testing.go:1146 +0x4b6
panic(0x11673c0, 0x128ffc0)
	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/panic.go:965 +0x1b9
github.com/karuppiah7890/redis-client-go/internal.(*MockRedisServer).Stop(0xc000060f28, 0x0, 0x0)
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:47 +0x2e
github.com/karuppiah7890/redis-client-go_test.TestConnect(0xc000186600)
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/client_test.go:24 +0x13f
testing.tRunner(0xc000186600, 0x119d690)
	/usr/local/Cellar/go/1.16.6/libexec/src/testing/testing.go:1193 +0xef
created by testing.(*T).Run
	/usr/local/Cellar/go/1.16.6/libexec/src/testing/testing.go:1238 +0x2b3
FAIL	github.com/karuppiah7890/redis-client-go	0.341s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

I didn't assign proper value to the `listener` field in `MockRedisServer`. My bad

I fixed it with one line `server.listener = l`

```go
func (server *MockRedisServer) Start() error {
	// Listen for incoming connections.
	l, err := net.Listen(server.server_type, fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		return fmt.Errorf("error listening for connections: %v", err.Error())
	}
	server.listener = l
	fmt.Printf("Listening at %s:%d\n", server.host, server.port)
	go func() {
		for {
			// Listen for an incoming connection.
			conn, err := l.Accept()
			if err != nil {
				fmt.Printf("error accepting connections: %v", err.Error())
			}
			// Handle connections in a new goroutine.
			go handleRequest(conn)
		}
	}()

	return nil
}
```

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
--- PASS: TestConnect (0.00s)
PASS
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionok  	github.com/karuppiah7890/redis-client-go	0.502s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

Lot of errors, hmm

But tests passed ;) Lol. Well. I just realized that the actual code doesn't do much. Hmm. 

I also need to write a better test which fails, which I can do only if I get some data from the `MockRedisServer` whether any connections were received, hmm

I added a field called `connections_received` to the mock server and did this - `server.connections_received++` -

```go
type MockRedisServer struct {
	host                 string
	port                 int
	server_type          string
	listener             net.Listener
	connections_received int
}

func (server *MockRedisServer) Start() error {
	// Listen for incoming connections.
	l, err := net.Listen(server.server_type, fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		return fmt.Errorf("error listening for connections: %v", err.Error())
	}
	server.listener = l
	fmt.Printf("Listening at %s:%d\n", server.host, server.port)
	go func() {
		for {
			// Listen for an incoming connection.
			conn, err := l.Accept()
			if err != nil {
				fmt.Printf("error accepting connections: %v", err.Error())
			}
			server.connections_received++
			// Handle connections in a new goroutine.
			go handleRequest(conn)
		}
	}()

	return nil
}

func (server *MockRedisServer) NumberOfConnectionsReceived() int {
	return server.connections_received;
}
```

And wrote this in the test -

```go
numberOfConnectionsReceived := server.NumberOfConnectionsReceived()

if numberOfConnectionsReceived != 1 {
    t.Errorf("Expected 1 connection to be received by the Redis Server but got: %v", numberOfConnectionsReceived)
}
```

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
    client_test.go:28: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection--- FAIL: TestConnect (0.00s)
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionerror accepting connections: accept tcp 127.0.0.1:6379: use of closed network connectionpanic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x11497a2]

goroutine 23 [running]:
github.com/karuppiah7890/redis-client-go/internal.handleRequest(0x0, 0x0)
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:63 +0x22
created by github.com/karuppiah7890/redis-client-go/internal.(*MockRedisServer).Start.func1
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:41 +0x51
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x11497a2]

goroutine 26 [running]:
github.com/karuppiah7890/redis-client-go/internal.handleRequest(0x0, 0x0)
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:63 +0x22
created by github.com/karuppiah7890/redis-client-go/internal.(*MockRedisServer).Start.func1
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:41 +0x51
FAIL	github.com/karuppiah7890/redis-client-go	0.296s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

Looks like there are some errors, hmm

```bash
goroutine 23 [running]:
github.com/karuppiah7890/redis-client-go/internal.handleRequest(0x0, 0x0)
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:63 +0x22
created by github.com/karuppiah7890/redis-client-go/internal.(*MockRedisServer).Start.func1
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:41 +0x51
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x11497a2]

goroutine 26 [running]:
github.com/karuppiah7890/redis-client-go/internal.handleRequest(0x0, 0x0)
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:63 +0x22
created by github.com/karuppiah7890/redis-client-go/internal.(*MockRedisServer).Start.func1
	/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:41 +0x51
```

Good thing is, test is also failing! :)

```bash
FAIL	github.com/karuppiah7890/redis-client-go	0.296s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
```

and it says 

`client_test.go:28: Expected 1 connection to be received by the Redis Server but got: 0`

But yeah, it's all too cluttered. Hmm

Let's look at the gorouting errors once. Lines 63 and 41

`/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:63`

`/Users/karuppiahn/projects/github.com/karuppiah7890/redis-client-go/internal/mock_redis_server.go:41`

Ahh. I found out the error. Hmm

The listen error passes through and since there's a `for` loop, it goes on and on and on again and again, hmm

The issue being here

```go
go func() {
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Printf("error accepting connections: %v", err.Error())
        }
        server.connections_received++
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}()
```

Note how the error is not caught and stopped

```go
if err != nil {
    fmt.Printf("error accepting connections: %v", err.Error())
}
```

It's weird that the connection count didn't increase though, hmm. I mean, the error didn't stop the processing. Maybe I read the value too early? I don't know. Anyways.

Let's fix this! :D

Oh. Now I understand what's going on. The `Accept` errors out because the test closes the server very fast and after closing, the `Accept` errors out saying - `error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection`

I just added another field called `running` in the Server. If the server isn't running, the listening loop will stop

```go
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
```

If there are errors in acception connections, I just print and continue. I also added new lines to the error logs

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
    client_test.go:28: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
--- FAIL: TestConnect (0.00s)
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
FAIL
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
FAIL	github.com/karuppiah7890/redis-client-go	0.362s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

Cool. Lot of errors, hmm

Oops. I forgot to assign `server.running = false` when `Stop()` is called

```go
func (server *MockRedisServer) Stop() error {
	server.running = false
	// Close the listener
	err := server.listener.Close()
	if err != nil {
		return fmt.Errorf("error closing listener: %v", err.Error())
	}
	return nil
}
```

Done

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
    client_test.go:28: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.343s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

There is still one error out there, hmm. I'm wondering how to do some graceful shutdown like stuff, hmm. Not easy. I did think about using channels, instead of all these fields. Hmm. Maybe this is okay for now. Hmm

The current server has been written in such a way that it can accept many connections from single or multiple clients. I guess this is okay and cool. However I'll be working with more tests soon. So. ;) But yeah, it's not gonna be easy to maintain this mock server. Lol. Hmm. I'm wondering how this is all gonna work. I could just an Actual Redis Server I guess. Hmm. I'm not sure how I'll verify if the client actually connected to it, hmm. I mean, I could simple check if there are no errors. Or just leave it. I mean...wait. This is kind of like me testing the `net` package which does `Dial`. RIGHT. LOL. Hmm. Anyways. I think I'm just going to try to get this over with and think more on the testing strategy. Hmm. Or else I might end up writing a Redis Server, in the name of a mock. I could use actual Redis Server too, and I did plan to use it anyway, like in integration tests. So yeah. Let's see. I guess even the current tests are like integration tests, but yeah, it's using a mock server, but it's still a running server, and checks the integration between the client and server, currently there's not much, but later, when the protocol is implemented, it would be testing quite some elaborate stuff, than just simple small units. Hmm

Cool. I finally implemented a basic connection in the client side

```go
package client

import (
	"fmt"
	"net"
)

func Connect(host string, port int) error {
	redisHost := fmt.Sprintf("%s:%v", host, port)

	conn, err := net.Dial("tcp", redisHost)
	if err != nil {
		return fmt.Errorf("error connecting to %s: %v", redisHost, err)
	}

	conn.Close()

	return err
}
```

and the test worked!!

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	0.378s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
```

yay! But yeah, I guess I did test the `net.Dial` feature. Lol. My bad. Hmm

Best thing is, this kind of test is faster than running a redis and running the golang test. Atleast that's my guess. Not saying that running redis server would take too long, but this is too fast. Surely with redis I would need to install redis server or use a redis server docker image and run it and also ensure ports are exposed for the Docker container, and then run the golang tests. So... ;) All good, I guess!

---

I was just reading the Redis Request Response model

https://redis.io/topics/protocol#request-response-model

And then the RESP protocol description

https://redis.io/topics/protocol#resp-protocol-description

I guess the following are pretty important -

```
For Simple Strings the first byte of the reply is "+"
For Errors the first byte of the reply is "-"
For Integers the first byte of the reply is ":"
For Bulk Strings the first byte of the reply is "$"
For Arrays the first byte of the reply is "*"
```

I also read the https://redis.io/topics/protocol#resp-simple-strings section

I'm planning to implement the `PING` command first, hmm. And run tests with an actual Redis server! :)

Maybe use Testcontainers to run redis ;)

https://duckduckgo.com/?t=ffab&q=testcontainers&ia=images

https://github.com/testcontainers/

I have used Testcontainers with Java before! - https://www.testcontainers.org/ , https://github.com/testcontainers/testcontainers-java

There's testcontainers for Golang too! :D

https://github.com/testcontainers/testcontainers-go

https://golang.testcontainers.org/

https://pkg.go.dev/github.com/testcontainers/testcontainers-go

https://golang.testcontainers.org/quickstart/gotest/

Perfect. There's a simple Redis example already

https://golang.testcontainers.org/quickstart/gotest/#2-spin-up-redis

I was just checking about the PING command

```bash
redis-client-go $ telnet localhost 6379
Trying ::1...
Connected to localhost.
Escape character is '^]'.
PING
+PONG
^]
telnet> Connection closed.
redis-client-go $ printf "PING" | nc localhost 6379
redis-client-go $ printf "PING\r\n" | nc localhost 6379
+PONG
redis-client-go $ 
```

Looks like I need to send `PING\r\n` and then look for `+PONG\r\n` and strip out the `+` and `\r\n` as mentioned in https://redis.io/topics/protocol#resp-simple-strings

Hmm. I'm wondering how to do this whole thing. Hmm. Maybe I'll start with something simple. Hmm

Just a `Ping` function

Damn. I was just changing some code. I was running actual Redis server locally, and I was running the tests and it kept failing and I was like "why isn't it working?" and it gave a weird error

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
    client_test.go:28: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.130s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
    client_test.go:28: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.097s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
    client_test.go:28: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.113s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:6379
error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	0.097s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

It should have said that it's not able to bind to the port 6379 as it's already in use. Instead it said `error accepting connections: accept tcp 127.0.0.1:6379: use of closed network connection`. Hmm. Weird.

Maybe I could use random ports ;)

`port := rand.Intn(65536)` 

Weird that it uses the same random number though. Hmm, like -

```bash
redis-client-go $ make test
go test -v ./...
^[[B=== RUN   TestConnect
Listening at localhost:33313
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

I tried to use `rand.Seed(10)` to simply see what it does and if it helps, I got a different result but it also failed my tests O.o

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:46190
    client_test.go:30: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:46190: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.343s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:46190
    client_test.go:30: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:46190: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.104s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:46190
    client_test.go:30: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:46190: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.100s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	0.382s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

It passed later when I commented it out. Hmm. I guess I could do away with some other port number or some other way of testing stuff, till I understand what's going on over here. Hmm. For now just a random port I guess

I think I need a way to ensure that a port is not open in my local before using it. Atleast in my bash, just to check if something else caused failures like the above

https://duckduckgo.com/?t=ffab&q=macos+open+port&ia=web&iax=qa

https://duckduckgo.com/?q=macos+open+port+check&t=ffab&ia=web

https://www.unixfu.ch/show-open-ports-on-mac/

```bash
lsof -i -P | grep -i "listen"
```

Cool! That works! I just tried to run a Redis server and use the above command and it worked!

```bash
$ lsof -i -P | grep -i "listen"
redis-ser 20925 karuppiahn    6u  IPv4 0x762d5b6bd440f2b5      0t0  TCP *:6379 (LISTEN)
redis-ser 20925 karuppiahn    7u  IPv6 0x762d5b6bcb84d18d      0t0  TCP *:6379 (LISTEN)
```

But it was a slow command too! Took a lot of time. Hmm

`ripgrep` / `rg` struggled too and took lot of time - like 3-5 seconds to run that command, hmm

```bash
$ lsof -i -P | rg -i "listen"
redis-ser 21656 karuppiahn    6u  IPv4 0x762d5b6bd3fb0705      0t0  TCP *:6379 (LISTEN)
redis-ser 21656 karuppiahn    7u  IPv6 0x762d5b6bcb84d18d      0t0  TCP *:6379 (LISTEN)
```
