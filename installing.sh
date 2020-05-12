brew install protobuf

protoc -I api/ -I ${GOPATH}/src \
   --go_out=plugins=grpc:api \
    api/api.proto

go build -i -v -o bin/server grpc/server
go build -i -v -o bin/client grpc/client