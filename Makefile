all: helloworld/helloworld.pb.go helloworld/helloworld.pb.gw.go helloworld/helloworld.swagger.json

helloworld/helloworld.pb.go: helloworld/helloworld.proto
		protoc -I/usr/local/include -I helloworld\
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/gengo/grpc-gateway/third_party/googleapis \
				--go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:helloworld \
				helloworld/helloworld.proto

helloworld/helloworld.pb.gw.go: helloworld/helloworld.proto
		protoc -I/usr/local/include -I helloworld\
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/gengo/grpc-gateway/third_party/googleapis \
				 --grpc-gateway_out=logtostderr=true:helloworld \
				helloworld/helloworld.proto

helloworld/helloworld.swagger.json: helloworld/helloworld.proto
		protoc -I/usr/local/include -I helloworld\
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/gengo/grpc-gateway/third_party/googleapis \
				--swagger_out=logtostderr=true:helloworld \
				helloworld/helloworld.proto
				
multiplica/multiplica.pb.go: multiplica/multiplica.proto
               protoc -IC:\2021\trabajonuevo\projectoos\multiplica\grpc\testgolangrpc\vendor\ -I=multiplica\ --go_out=multiplica\ multiplica\multiplica.proto
