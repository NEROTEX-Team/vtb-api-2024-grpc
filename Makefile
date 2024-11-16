
generate-proto:
	protoc --go_out=. --go-grpc_out=. api/v1/user/user.proto

update-packages:
	go mod tidy && go mod vendor

build:
	go build -o server.bin cmd/grpc_server/main.go

local:
	docker compose -f docker-compose.dev.yaml up --force-recreate --renew-anon-volumes --build

local-down:
	docker compose -f docker-compose.dev.yaml down -v