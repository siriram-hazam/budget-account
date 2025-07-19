.PHONY: build

build:
	go build -o main ./cmd/auth

run:
	go run ./cmd/auth

test:
	go test ./...

migrate:
	migrate -path db/migrations -database "postgres://postgres:123456@localhost:5432/auth?sslmode=disable" up

proto:
	protoc --go_out=grpc-auth/proto --go-grpc_out=grpc-auth/proto proto/auth.proto

air:
	air

docker-build:
	docker build -t budget-authen:latest .

docker-run:
	docker run --rm -p 8080:8080 budget-authen:latest

docker-dev:
    docker run --rm -p 8080:8080 -v $(PWD):/app --entrypoint air budget-authen:latest