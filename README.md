# GRPC on Heroku

A simple GRPC service, written in Go, running on Heroku. Since Heroku doesn't
support HTTP/2, this service is exposed as an HTTP+JSON API.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

protoc -IC:\2021\trabajonuevo\projectoos\multiplica\grpc\testgolangrpc\vendor\ -I=multiplica\ --go_out=multiplica\ multiplica\multiplica.proto

