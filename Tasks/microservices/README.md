## Now only one main file in repo
it accepts 1 or 3 command line arguments

if you wanna to start a server (not proxy):
```
go run cmd/main.go <port>
```

if you want to start a proxy with servers then:
```
go run cmd/main.go <port server 1> <port server 2> <proxy port>
```

### You can start many instances of app with proxy:
but you have to change the package of the cmd/autoSpammer to main

```
go run cmd/autoSpammer/mainManyInstances.go <number of servers>
```

It will start the amount of servers and proxy starting from port 60490

### To start up the database:
```
docker-compose up
```

