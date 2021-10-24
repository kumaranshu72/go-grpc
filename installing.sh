// installing proto complier
brew install protobuf

ls

// compling ptoyobuf
protoc -I api/ -I ${GOPATH}/src \
   --go_out=plugins=grpc:api \
    api/api.proto

// building server & client
go build -i -v -o bin/server grpc/server
go build -i -v -o bin/client grpc/client

// generating Self signed ssl certificate
openssl genrsa -out cert/server.key 2048
openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
openssl req -new -sha256 -key cert/server.key -out cert/server.csr
openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650


