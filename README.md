#  overview
study grpc using golang

# wok procedure

## 1. build docker

```shell script
docker-compose up
``` 

## 2. run server.go & run client.go
```shell script
# Enter docker container
docek exec -it sample-grpc bash

# run server & run client.go
go run server.go || go rnu client.go
# > result:name:"mike" kind:"Norwegian Forest Cat" 
```