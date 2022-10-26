### You can start many instances of app with proxy:

```go run cmd/mainManyInstances.go <number of servers>```

It will start the amount of servers and proxy

### To start up the database:
```docker-compose up```

### To start the services separately need to run the following commands:

```go run cmd/app/first.go```

```go run cmd/app/second.go```

```go run cmd/app/proxy.go```