
generate_proto:
	protoc --go_out=. --go-grpc_out=. api/v1/user/user.proto

update_packages:
	go mod tidy && go mod vendor

build:
	go build -o server.bin cmd/grpc_server/main.go
