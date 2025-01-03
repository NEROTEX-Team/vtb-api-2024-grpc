
generate-proto:
	protoc --go_out=. --go-grpc_out=. api/v1/user/user.proto

update-packages:
	go mod tidy && go mod vendor

build-local:
	go build -o server.bin cmd/grpc_server/main.go

build-docker:
	docker build . -t vtb-api-2024-grpc-server
	cd client && docker build . -t vtb-api-2024-grpc-client

test:
	go test -v ./...

local:
	docker compose -f docker-compose.dev.yaml up --force-recreate --renew-anon-volumes --build

local-down:
	docker compose -f docker-compose.dev.yaml down -v

cert:
	cd ./certs; bash gen-cert.sh; cd ..