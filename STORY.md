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

---

About `rand`, previously I was trying `math/rand`. I didn't try `crypto/rand`, so..


https://pkg.go.dev/crypto/rand?utm_source=gopls

https://pkg.go.dev/crypto/rand?utm_source=gopls#Int

But that returns `big.Int`. Right. Actually `*big.Int`. Hmm

`port := rand.Int(rand.Reader, big.NewInt(65536))`

I needed to also handle error

Meh. Gonna skip that. It's okay to use some random port number, maybe something other than 6379 for now. Hmm. Or let it be 6379, as testcontainers creates random ports and exposes them, so wohoo! :) I do need to ensure that redis isn't running locally with default port, hmm. For now just using `port := rand.Intn(65536)` !

---

Moving on to the PING command! :D

I was trying to get rid of the `conn.Close()` in the `Connect` function as that closes the connection just after connecting to the server. But we need to communicate with the server later. When I removed this, the tests fail, hmm, weird

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:29: Expected 1 connection to be received by the Redis Server but got: 0
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.378s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

Looks like a weird thing. Makes me wonder if I should get rid of the test I have and just use actual redis server. Hmm. Something for me to learn and understand to see why these random errors are popping up. I just didn't close the connection, but I did connect, hmm

Wow, there are failures when I put close too, at times. The close is in the test now!

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- PASS: TestConnect (0.00s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	0.100s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:31: Expected 1 connection to be received by the Redis Server but got: 0
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.386s
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
ok  	github.com/karuppiah7890/redis-client-go	0.099s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
```

It looks like this -

```go
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
```

If I close after checking connections, it fails

```go
numberOfConnectionsReceived := server.NumberOfConnectionsReceived()
conn.Close()
```

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:30: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.366s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:30: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- FAIL: TestConnect (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.187s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
redis-client-go $ lsof -i -P | grep -i "listen"
```

I think the accept of connection is not happening properly, or basically the dial ain't seem to be happening properly. Hmm

I'm just gonna close it before hand, and then check the number of connections in mock redis server :P

---

I'm trying to work on the PING command and I want to run an actual redis for it. So for testcontainers, I'm doing this -

https://golang.testcontainers.org/quickstart/gotest/

```bash
$ go get -u github.com/testcontainers/testcontainers-go
go: downloading github.com/testcontainers/testcontainers-go v0.11.1
go: downloading github.com/docker/docker v20.10.7+incompatible
go: downloading github.com/google/uuid v1.3.0
go: downloading github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6
go: downloading golang.org/x/sys v0.0.0-20210820121016-41cdb8703e55
go: downloading github.com/docker/docker v20.10.8+incompatible
go: downloading github.com/containerd/containerd v1.5.0-beta.4
go: downloading github.com/containerd/containerd v1.5.5
go: downloading github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1
go: downloading github.com/Microsoft/hcsshim v0.8.16
go: downloading github.com/Microsoft/go-winio v0.5.0
go: downloading github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c
go: downloading google.golang.org/grpc v1.33.2
go: downloading golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
go: downloading github.com/Microsoft/hcsshim v0.8.20
go: downloading google.golang.org/grpc v1.40.0
go: downloading google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a
go: downloading github.com/opencontainers/runc v1.0.1
go: downloading go.opencensus.io v0.22.3
go: downloading google.golang.org/genproto v0.0.0-20210821163610-241b8fcbd6c8
go: downloading github.com/moby/sys/mount v0.2.0
go: downloading github.com/moby/sys v0.0.0-20210813220516-f3885c897d0f
go: downloading github.com/containerd/cgroups v0.0.0-20210114181951-8a68de567b68
go: downloading github.com/containerd/cgroups v1.0.1
go: downloading github.com/cenkalti/backoff v1.1.0
go get: added github.com/Microsoft/go-winio v0.5.0
go get: added github.com/Microsoft/hcsshim v0.8.20
go get: added github.com/containerd/containerd v1.5.5
go get: added github.com/docker/docker v20.10.8+incompatible
go get: added github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
go get: added github.com/google/uuid v1.3.0
go get: added github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6
go get: added github.com/morikuni/aec v1.0.0
go get: added github.com/testcontainers/testcontainers-go v0.11.1
go get: added go.opencensus.io v0.23.0
go get: added golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
go get: added golang.org/x/sys v0.0.0-20210820121016-41cdb8703e55
go get: added google.golang.org/genproto v0.0.0-20210821163610-241b8fcbd6c8
```

```bash
redis-client-go $ go mod tidy -v
go: downloading github.com/go-redis/redis v6.15.9+incompatible
go: downloading github.com/go-sql-driver/mysql v1.6.0
go: downloading gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
go: downloading github.com/containerd/continuity v0.1.0
go: downloading github.com/gorilla/mux v1.7.2
go: downloading github.com/kr/pretty v0.2.1
redis-client-go $ go mod tidy -v
redis-client-go $ 
```

```go
package internal

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartRedisServer(t *testing.T) {
	ctx := context.Background()
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
		t.Error(err)
	}
	defer redisC.Terminate(ctx)
}
```

Changing this to -

```go
package internal

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartRedisServer(ctx context.Context) (testcontainers.Container, error) {
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
		return nil, err
	}
	return redisC, nil
}
```

and the test file -

```go
func TestPing(t *testing.T) {
	ctx := context.Background()
	redisC, err := internal.StartRedisServer(ctx)

	if err != nil {
		t.Errorf("failed to start the redis container: %v", err)
		return
	}

	defer redisC.Terminate(ctx)
}
```

I just tried it out.

I'm wondering if anything happened at all in later runs

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
--- PASS: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 00:31:35 Starting container id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:35 Waiting for container id 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:36 Container is ready id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:49 Starting container id: 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Waiting for container id 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Container is ready id: 35f5b0f67ccc image: redis:latest
--- PASS: TestPing (22.11s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	22.590s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
--- PASS: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 00:31:35 Starting container id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:35 Waiting for container id 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:36 Container is ready id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:49 Starting container id: 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Waiting for container id 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Container is ready id: 35f5b0f67ccc image: redis:latest
--- PASS: TestPing (22.11s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
--- PASS: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 00:31:35 Starting container id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:35 Waiting for container id 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:36 Container is ready id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:49 Starting container id: 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Waiting for container id 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Container is ready id: 35f5b0f67ccc image: redis:latest
--- PASS: TestPing (22.11s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
--- PASS: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 00:31:35 Starting container id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:35 Waiting for container id 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:36 Container is ready id: 1408d446803b image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:31:49 Starting container id: 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Waiting for container id 35f5b0f67ccc image: redis:latest
2021/08/22 00:31:49 Container is ready id: 35f5b0f67ccc image: redis:latest
--- PASS: TestPing (22.11s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

The time and log etc is all the same. No new docker containers later. Oh. I just noticed it is all `(cached)`. hmm

No wonder this is what I saw previously too. In my logs. All those usage of same port. It was just run once, and the same logs were shown. Right. Nice. `(cached)`

I was checking if I can disable cache to just see the speed. Hmm

```bash
redis-client-go $ go test -h
usage: go test [build/test flags] [packages] [build/test flags & test binary flags]
Run 'go help test' and 'go help testflag' for details.
redis-client-go $ go help testflag
The 'go test' command takes both flags that apply to 'go test' itself
and flags that apply to the resulting test binary.

Several of the flags control profiling and write an execution profile
suitable for "go tool pprof"; run "go tool pprof -h" for more
information. The --alloc_space, --alloc_objects, and --show_bytes
options of pprof control how the information is presented.

The following flags are recognized by the 'go test' command and
control the execution of any test:

	-bench regexp
	    Run only those benchmarks matching a regular expression.
	    By default, no benchmarks are run.
	    To run all benchmarks, use '-bench .' or '-bench=.'.
	    The regular expression is split by unbracketed slash (/)
	    characters into a sequence of regular expressions, and each
	    part of a benchmark's identifier must match the corresponding
	    element in the sequence, if any. Possible parents of matches
	    are run with b.N=1 to identify sub-benchmarks. For example,
	    given -bench=X/Y, top-level benchmarks matching X are run
	    with b.N=1 to find any sub-benchmarks matching Y, which are
	    then run in full.

	-benchtime t
	    Run enough iterations of each benchmark to take t, specified
	    as a time.Duration (for example, -benchtime 1h30s).
	    The default is 1 second (1s).
	    The special syntax Nx means to run the benchmark N times
	    (for example, -benchtime 100x).

	-count n
	    Run each test and benchmark n times (default 1).
	    If -cpu is set, run n times for each GOMAXPROCS value.
	    Examples are always run once.

	-cover
	    Enable coverage analysis.
	    Note that because coverage works by annotating the source
	    code before compilation, compilation and test failures with
	    coverage enabled may report line numbers that don't correspond
	    to the original sources.

	-covermode set,count,atomic
	    Set the mode for coverage analysis for the package[s]
	    being tested. The default is "set" unless -race is enabled,
	    in which case it is "atomic".
	    The values:
		set: bool: does this statement run?
		count: int: how many times does this statement run?
		atomic: int: count, but correct in multithreaded tests;
			significantly more expensive.
	    Sets -cover.

	-coverpkg pattern1,pattern2,pattern3
	    Apply coverage analysis in each test to packages matching the patterns.
	    The default is for each test to analyze only the package being tested.
	    See 'go help packages' for a description of package patterns.
	    Sets -cover.

	-cpu 1,2,4
	    Specify a list of GOMAXPROCS values for which the tests or
	    benchmarks should be executed. The default is the current value
	    of GOMAXPROCS.

	-failfast
	    Do not start new tests after the first test failure.

	-list regexp
	    List tests, benchmarks, or examples matching the regular expression.
	    No tests, benchmarks or examples will be run. This will only
	    list top-level tests. No subtest or subbenchmarks will be shown.

	-parallel n
	    Allow parallel execution of test functions that call t.Parallel.
	    The value of this flag is the maximum number of tests to run
	    simultaneously; by default, it is set to the value of GOMAXPROCS.
	    Note that -parallel only applies within a single test binary.
	    The 'go test' command may run tests for different packages
	    in parallel as well, according to the setting of the -p flag
	    (see 'go help build').

	-run regexp
	    Run only those tests and examples matching the regular expression.
	    For tests, the regular expression is split by unbracketed slash (/)
	    characters into a sequence of regular expressions, and each part
	    of a test's identifier must match the corresponding element in
	    the sequence, if any. Note that possible parents of matches are
	    run too, so that -run=X/Y matches and runs and reports the result
	    of all tests matching X, even those without sub-tests matching Y,
	    because it must run them to look for those sub-tests.

	-short
	    Tell long-running tests to shorten their run time.
	    It is off by default but set during all.bash so that installing
	    the Go tree can run a sanity check but not spend time running
	    exhaustive tests.

	-timeout d
	    If a test binary runs longer than duration d, panic.
	    If d is 0, the timeout is disabled.
	    The default is 10 minutes (10m).

	-v
	    Verbose output: log all tests as they are run. Also print all
	    text from Log and Logf calls even if the test succeeds.

	-vet list
	    Configure the invocation of "go vet" during "go test"
	    to use the comma-separated list of vet checks.
	    If list is empty, "go test" runs "go vet" with a curated list of
	    checks believed to be always worth addressing.
	    If list is "off", "go test" does not run "go vet" at all.

The following flags are also recognized by 'go test' and can be used to
profile the tests during execution:

	-benchmem
	    Print memory allocation statistics for benchmarks.

	-blockprofile block.out
	    Write a goroutine blocking profile to the specified file
	    when all tests are complete.
	    Writes test binary as -c would.

	-blockprofilerate n
	    Control the detail provided in goroutine blocking profiles by
	    calling runtime.SetBlockProfileRate with n.
	    See 'go doc runtime.SetBlockProfileRate'.
	    The profiler aims to sample, on average, one blocking event every
	    n nanoseconds the program spends blocked. By default,
	    if -test.blockprofile is set without this flag, all blocking events
	    are recorded, equivalent to -test.blockprofilerate=1.

	-coverprofile cover.out
	    Write a coverage profile to the file after all tests have passed.
	    Sets -cover.

	-cpuprofile cpu.out
	    Write a CPU profile to the specified file before exiting.
	    Writes test binary as -c would.

	-memprofile mem.out
	    Write an allocation profile to the file after all tests have passed.
	    Writes test binary as -c would.

	-memprofilerate n
	    Enable more precise (and expensive) memory allocation profiles by
	    setting runtime.MemProfileRate. See 'go doc runtime.MemProfileRate'.
	    To profile all memory allocations, use -test.memprofilerate=1.

	-mutexprofile mutex.out
	    Write a mutex contention profile to the specified file
	    when all tests are complete.
	    Writes test binary as -c would.

	-mutexprofilefraction n
	    Sample 1 in n stack traces of goroutines holding a
	    contended mutex.

	-outputdir directory
	    Place output files from profiling in the specified directory,
	    by default the directory in which "go test" is running.

	-trace trace.out
	    Write an execution trace to the specified file before exiting.

Each of these flags is also recognized with an optional 'test.' prefix,
as in -test.v. When invoking the generated test binary (the result of
'go test -c') directly, however, the prefix is mandatory.

The 'go test' command rewrites or removes recognized flags,
as appropriate, both before and after the optional package list,
before invoking the test binary.

For instance, the command

	go test -v -myflag testdata -cpuprofile=prof.out -x

will compile the test binary and then run it as

	pkg.test -test.v -myflag testdata -test.cpuprofile=prof.out

(The -x flag is removed because it applies only to the go command's
execution, not to the test itself.)

The test flags that generate profiles (other than for coverage) also
leave the test binary in pkg.test for use when analyzing the profiles.

When 'go test' runs a test binary, it does so from within the
corresponding package's source code directory. Depending on the test,
it may be necessary to do the same when invoking a generated test
binary directly.

The command-line package list, if present, must appear before any
flag not known to the go test command. Continuing the example above,
the package list would have to appear before -myflag, but could appear
on either side of -v.

When 'go test' runs in package list mode, 'go test' caches successful
package test results to avoid unnecessary repeated running of tests. To
disable test caching, use any test flag or argument other than the
cacheable flags. The idiomatic way to disable test caching explicitly
is to use -count=1.

To keep an argument for a test binary from being interpreted as a
known flag or a package name, use -args (see 'go help test') which
passes the remainder of the command line through to the test binary
uninterpreted and unaltered.

For instance, the command

	go test -v -args -x -v

will compile the test binary and then run it as

	pkg.test -test.v -x -v

Similarly,

	go test -args math

will compile the test binary and then run it as

	pkg.test math

In the first example, the -x and the second -v are passed through to the
test binary unchanged and with no effect on the go command itself.
In the second example, the argument math is passed through to the test
binary, instead of being interpreted as the package list.
redis-client-go $ go help test
usage: go test [build/test flags] [packages] [build/test flags & test binary flags]

'Go test' automates testing the packages named by the import paths.
It prints a summary of the test results in the format:

	ok   archive/tar   0.011s
	FAIL archive/zip   0.022s
	ok   compress/gzip 0.033s
	...

followed by detailed output for each failed package.

'Go test' recompiles each package along with any files with names matching
the file pattern "*_test.go".
These additional files can contain test functions, benchmark functions, and
example functions. See 'go help testfunc' for more.
Each listed package causes the execution of a separate test binary.
Files whose names begin with "_" (including "_test.go") or "." are ignored.

Test files that declare a package with the suffix "_test" will be compiled as a
separate package, and then linked and run with the main test binary.

The go tool will ignore a directory named "testdata", making it available
to hold ancillary data needed by the tests.

As part of building a test binary, go test runs go vet on the package
and its test source files to identify significant problems. If go vet
finds any problems, go test reports those and does not run the test
binary. Only a high-confidence subset of the default go vet checks are
used. That subset is: 'atomic', 'bool', 'buildtags', 'errorsas',
'ifaceassert', 'nilfunc', 'printf', and 'stringintconv'. You can see
the documentation for these and other vet tests via "go doc cmd/vet".
To disable the running of go vet, use the -vet=off flag.

All test output and summary lines are printed to the go command's
standard output, even if the test printed them to its own standard
error. (The go command's standard error is reserved for printing
errors building the tests.)

Go test runs in two different modes:

The first, called local directory mode, occurs when go test is
invoked with no package arguments (for example, 'go test' or 'go
test -v'). In this mode, go test compiles the package sources and
tests found in the current directory and then runs the resulting
test binary. In this mode, caching (discussed below) is disabled.
After the package test finishes, go test prints a summary line
showing the test status ('ok' or 'FAIL'), package name, and elapsed
time.

The second, called package list mode, occurs when go test is invoked
with explicit package arguments (for example 'go test math', 'go
test ./...', and even 'go test .'). In this mode, go test compiles
and tests each of the packages listed on the command line. If a
package test passes, go test prints only the final 'ok' summary
line. If a package test fails, go test prints the full test output.
If invoked with the -bench or -v flag, go test prints the full
output even for passing package tests, in order to display the
requested benchmark results or verbose logging. After the package
tests for all of the listed packages finish, and their output is
printed, go test prints a final 'FAIL' status if any package test
has failed.

In package list mode only, go test caches successful package test
results to avoid unnecessary repeated running of tests. When the
result of a test can be recovered from the cache, go test will
redisplay the previous output instead of running the test binary
again. When this happens, go test prints '(cached)' in place of the
elapsed time in the summary line.

The rule for a match in the cache is that the run involves the same
test binary and the flags on the command line come entirely from a
restricted set of 'cacheable' test flags, defined as -cpu, -list,
-parallel, -run, -short, and -v. If a run of go test has any test
or non-test flags outside this set, the result is not cached. To
disable test caching, use any test flag or argument other than the
cacheable flags. The idiomatic way to disable test caching explicitly
is to use -count=1. Tests that open files within the package's source
root (usually $GOPATH) or that consult environment variables only
match future runs in which the files and environment variables are unchanged.
A cached test result is treated as executing in no time at all,
so a successful package test result will be cached and reused
regardless of -timeout setting.

In addition to the build flags, the flags handled by 'go test' itself are:

	-args
	    Pass the remainder of the command line (everything after -args)
	    to the test binary, uninterpreted and unchanged.
	    Because this flag consumes the remainder of the command line,
	    the package list (if present) must appear before this flag.

	-c
	    Compile the test binary to pkg.test but do not run it
	    (where pkg is the last element of the package's import path).
	    The file name can be changed with the -o flag.

	-exec xprog
	    Run the test binary using xprog. The behavior is the same as
	    in 'go run'. See 'go help run' for details.

	-i
	    Install packages that are dependencies of the test.
	    Do not run the test.
	    The -i flag is deprecated. Compiled packages are cached automatically.

	-json
	    Convert test output to JSON suitable for automated processing.
	    See 'go doc test2json' for the encoding details.

	-o file
	    Compile the test binary to the named file.
	    The test still runs (unless -c or -i is specified).

The test binary also accepts flags that control execution of the test; these
flags are also accessible by 'go test'. See 'go help testflag' for details.

For more about build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go vet.
redis-client-go $ 
```

https://duckduckgo.com/?t=ffab&q=go+test+disable+cache&ia=web

https://til.cybertec-postgresql.com/post/2019-11-07-How-to-turn-off-test-caching-for-golang/

```bash
go clean -testcache
```

There's one more for build cache.

Oh wow. We need to do quite a lot to clean up the cache huh. Not just a simple flag in test command huh. Hmm. Okay. Meh. Surely I'll change code. I'll see it run anew and notice the time. Hopefully faster the next time. But maybe not too fast, unfortunately.

I was seeing a lot of passing tests. I was wondering if everything was okay :P

I wanted to see how a failing test looks like and tried with a different redis port to connect to, to see if it failed. It did fail. But weirdly both ping and connect tests failed

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:31: Expected 1 connection to be received by the Redis Server but got: 0
--- FAIL: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 00:53:37 Starting container id: f2694489d571 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:53:38 Waiting for container id f2694489d571 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:53:38 Container is ready id: f2694489d571 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:53:38 Starting container id: 7062d56d7345 image: redis:latest
2021/08/22 00:53:38 Waiting for container id 7062d56d7345 image: redis:latest
2021/08/22 00:53:38 Container is ready id: 7062d56d7345 image: redis:latest
    client_test.go:61: Connection to Redis Server failed: error connecting to localhost:58753: dial tcp [::1]:58753: connect: connection refused
--- FAIL: TestPing (1.35s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	1.716s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
--- PASS: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 00:52:38 Starting container id: f1f3d2b0c476 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:52:39 Waiting for container id f1f3d2b0c476 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:52:39 Container is ready id: f1f3d2b0c476 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 00:52:39 Starting container id: f9e5dfeff2af image: redis:latest
2021/08/22 00:52:39 Waiting for container id f9e5dfeff2af image: redis:latest
2021/08/22 00:52:39 Container is ready id: f9e5dfeff2af image: redis:latest
--- PASS: TestPing (1.97s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

`TestConnect` is some weird test. Hmm. I think it's flaky and has timing issues. Hmm

`TestPing` looks like this now

```go
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
```

---

I'm still implementing `PING`. I was just wondering how Errors work and was checking how redis-cli works, like, if it gives errors on client side for wrong commands or if the command is sent to server and server replies back and tells that the command is wrong etc. Looks like the command is sent to server. I guess it makes sense because server only has the feature. Also, servers have Redis module feature too, I think that allows dynamic commands to be added? I don't know. So it's hard for clients to maybe find out about those dynamic commands at the time of execution, so sending and asking server to give a reply is better

```bash
~ $ printf "BLAH\r\n" | nc localhost 6379
-ERR unknown command `BLAH`, with args beginning with: 
```

```bash
~ $ redis-cli
127.0.0.1:6379> BLAH
(error) ERR unknown command `BLAH`, with args beginning with: 
127.0.0.1:6379>
~ $ 
```

---

About the flaky failures, I was thinking about it. This is what I was guessing -

I think connecting / dialling and immediately closing the connection errors things on server side when things are slow, like, I think as part of accept, server does lot of stuff and then returns connection and error if any. I think sometimes when we close too early, it's so early that accept is not done and that's when we get the connection closed error I think. Just a guess. If we just write to the connection and read something from the server and then close the connection, maybe we wouldn't see any such errors, because we would write to the connection and then wait for reading from server and then only close the connection, which is surely way beyond accept function in server. Hmm

---

I just did some small refactoring for running the redis in testcontainers and also get the host and port in the same function

Too many errors, hmm

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:31: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- FAIL: TestConnect (0.00s)
=== RUN   TestPing
    client_test.go:39: failed to start the redis container: failed to create container: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
--- FAIL: TestPing (0.00s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	0.404s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:31: Expected 1 connection to be received by the Redis Server but got: 0
--- FAIL: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:22:44 Starting container id: 316dc3bd39c9 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:22:44 Waiting for container id 316dc3bd39c9 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:22:44 Container is ready id: 316dc3bd39c9 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:22:45 Starting container id: 426e720605af image: redis:latest
2021/08/22 11:22:45 Waiting for container id 426e720605af image: redis:latest
2021/08/22 11:22:45 Container is ready id: 426e720605af image: redis:latest
--- PASS: TestPing (1.12s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	1.295s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:31: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- FAIL: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:22:52 Starting container id: 765e99070f1d image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:22:52 Waiting for container id 765e99070f1d image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:22:52 Container is ready id: 765e99070f1d image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:22:52 Starting container id: b2da5fbd188b image: redis:latest
2021/08/22 11:22:52 Waiting for container id b2da5fbd188b image: redis:latest
2021/08/22 11:22:52 Container is ready id: b2da5fbd188b image: redis:latest
--- PASS: TestPing (0.99s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	1.112s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
Listening at localhost:33313
    client_test.go:31: Expected 1 connection to be received by the Redis Server but got: 0
error accepting connections: accept tcp 127.0.0.1:33313: use of closed network connection
--- FAIL: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:23:01 Starting container id: 1396dfcf2961 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:23:01 Waiting for container id 1396dfcf2961 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:23:01 Container is ready id: 1396dfcf2961 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:23:01 Starting container id: 156d34028ff9 image: redis:latest
2021/08/22 11:23:01 Waiting for container id 156d34028ff9 image: redis:latest
2021/08/22 11:23:01 Container is ready id: 156d34028ff9 image: redis:latest
--- PASS: TestPing (0.96s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	1.086s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

I'm planning to skip the `TestConnect` test for now. Hmm

Done!

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
    client_test.go:13: 
--- SKIP: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:25:00 Starting container id: e21474db2ed6 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:25:00 Waiting for container id e21474db2ed6 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:25:00 Container is ready id: e21474db2ed6 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:25:00 Starting container id: 53cda0cf887c image: redis:latest
2021/08/22 11:25:01 Waiting for container id 53cda0cf887c image: redis:latest
2021/08/22 11:25:01 Container is ready id: 53cda0cf887c image: redis:latest
--- PASS: TestPing (0.91s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	1.305s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
    client_test.go:13: 
--- SKIP: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:25:00 Starting container id: e21474db2ed6 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:25:00 Waiting for container id e21474db2ed6 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:25:00 Container is ready id: e21474db2ed6 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:25:00 Starting container id: 53cda0cf887c image: redis:latest
2021/08/22 11:25:01 Waiting for container id 53cda0cf887c image: redis:latest
2021/08/22 11:25:01 Container is ready id: 53cda0cf887c image: redis:latest
--- PASS: TestPing (0.91s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	(cached)
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ gst
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   client_test.go

no changes added to commit (use "git add" and/or "git commit -a")
redis-client-go $ ga .
redis-client-go $ gc -m "skip flaky TestConnect test"
[main a884ffa] skip flaky TestConnect test
 1 file changed, 1 insertion(+)
redis-client-go $ 
```

---

Cool, now I wrote a test for PING that fails!

```bash
redis-client-go $ gst
On branch main
Your branch is ahead of 'origin/main' by 2 commits.
  (use "git push" to publish your local commits)

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   client_test.go

no changes added to commit (use "git add" and/or "git commit -a")
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
    client_test.go:13: 
--- SKIP: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:29:49 Starting container id: 3bcf0f4179de image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:29:49 Waiting for container id 3bcf0f4179de image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:29:49 Container is ready id: 3bcf0f4179de image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:29:49 Starting container id: 4c0bf8239fbd image: redis:latest
2021/08/22 11:29:50 Waiting for container id 4c0bf8239fbd image: redis:latest
2021/08/22 11:29:50 Container is ready id: 4c0bf8239fbd image: redis:latest
--- PASS: TestPing (0.96s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	1.401s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
    client_test.go:13: 
--- SKIP: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:30:18 Starting container id: 1d136a26317a image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:30:18 Waiting for container id 1d136a26317a image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:30:18 Container is ready id: 1d136a26317a image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:30:18 Starting container id: d219e9728060 image: redis:latest
2021/08/22 11:30:18 Waiting for container id d219e9728060 image: redis:latest
2021/08/22 11:30:18 Container is ready id: d219e9728060 image: redis:latest
    client_test.go:60: Expected PONG as reply for PING but got: 
--- FAIL: TestPing (0.94s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	1.258s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

```go
	pingResponse, err := client.Ping(conn)
	if err != nil {
		t.Errorf("Expected no errors in PING but got: %v", err)
	}

	if pingResponse != "PONG" {
		t.Errorf("Expected PONG as reply for PING but got: %v", pingResponse)
	}
```

Initially I made a mistake of writing `if pingResponse == "PONG" {` lol and test was passing and I was like "what? I didn't implement anything yet" and then noticed the issue in the test :P

Now I'm trying to implement it. For it I'm using `conn.Write`

https://pkg.go.dev/net#Conn.Write

I'm guessing the return value `n` is to tell how much bytes has been written, hmm

---

I implemented ping and got errors in test

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
    client_test.go:13: 
--- SKIP: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:43:46 Starting container id: 99ffb2b9d8f7 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:43:47 Waiting for container id 99ffb2b9d8f7 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:43:47 Container is ready id: 99ffb2b9d8f7 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:43:47 Starting container id: 0e5750f07f02 image: redis:latest
2021/08/22 11:43:47 Waiting for container id 0e5750f07f02 image: redis:latest
2021/08/22 11:43:47 Container is ready id: 0e5750f07f02 image: redis:latest
    client_test.go:56: Expected no errors in PING but got: error while pinging: write tcp [::1]:49860->[::1]:49859: use of closed network connection
    client_test.go:60: Expected PONG as reply for PING but got: 
--- FAIL: TestPing (0.96s)
FAIL
FAIL	github.com/karuppiah7890/redis-client-go	1.382s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
FAIL
make: *** [test] Error 1
redis-client-go $ 
```

Okay. Looks like I closed the connection before doing ping, lol. Missed the `defer`

```go
conn, err := client.Connect(host, port)
if err != nil {
	t.Errorf("Connection to Redis Server failed: %v", err)
	return
}

conn.Close()

pingResponse, err := client.Ping(conn)
if err != nil {
	t.Errorf("Expected no errors in PING but got: %v", err)
}

if pingResponse != "PONG" {
	t.Errorf("Expected PONG as reply for PING but got: %v", pingResponse)
}
```

Note the `conn.Close`. Lol. The `defer`rrrr

And with the `defer`, everything passes!! :D yay

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
    client_test.go:13: 
--- SKIP: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:45:15 Starting container id: 10c3bdab2bb5 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:45:15 Waiting for container id 10c3bdab2bb5 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:45:15 Container is ready id: 10c3bdab2bb5 image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:45:15 Starting container id: 113d539a38ca image: redis:latest
2021/08/22 11:45:16 Waiting for container id 113d539a38ca image: redis:latest
2021/08/22 11:45:16 Container is ready id: 113d539a38ca image: redis:latest
--- PASS: TestPing (0.94s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	1.330s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

The implementation is still too meh and crazy actually. Too many error checks and a bit of a weird implementation, hmm

```go
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
```

I put a println statement in between to check how the whole thing looks like

`fmt.Println(string(buf))`

```bash
redis-client-go $ make test
go test -v ./...
=== RUN   TestConnect
    client_test.go:13: 
--- SKIP: TestConnect (0.00s)
=== RUN   TestPing
2021/08/22 11:46:24 Starting container id: 0611e121219d image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:46:25 Waiting for container id 0611e121219d image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:46:25 Container is ready id: 0611e121219d image: quay.io/testcontainers/ryuk:0.2.3
2021/08/22 11:46:25 Starting container id: 9a4e6df922cf image: redis:latest
2021/08/22 11:46:25 Waiting for container id 9a4e6df922cf image: redis:latest
2021/08/22 11:46:25 Container is ready id: 9a4e6df922cf image: redis:latest
+PONG

--- PASS: TestPing (1.68s)
PASS
ok  	github.com/karuppiah7890/redis-client-go	2.058s
?   	github.com/karuppiah7890/redis-client-go/internal	[no test files]
redis-client-go $ 
```

It does say `+PONG` :D
