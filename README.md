# GRPC on Heroku

A simple GRPC service, written in Go, running on Heroku. Since Heroku doesn't
support HTTP/2, this service is exposed as an HTTP+JSON API.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

protoc -IC:\2021\trabajonuevo\projectoos\multiplica\grpc\testgolangrpc\vendor\ -I=multiplica\ --go_out=multiplica\ multiplica\multiplica.proto





example:

http://youtochigrpc.herokuapp.com/v2/multiplica/greeter/say-multiplica
method POST
using Postman:   raw json
{
   "numero" :"18",
    "veces" :"20"
}

result
{
    "message": "360.000000"
}


example:

http://youtochigrpc.herokuapp.com/v2/multiplica/greeter/say-multiplica
method POST
using Postman:   raw json

{
   "numero" :"1",
    "veces" :"2"
}

result
{
    "message": "2.000000"
}
