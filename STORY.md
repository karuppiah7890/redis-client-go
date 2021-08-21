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

