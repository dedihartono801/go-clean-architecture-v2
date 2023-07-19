GOPATH ?= $(HOME)/go

ifeq ($(OS), Windows_NT)
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
else
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
endif

run-rest:
	go run cmd/http/main.go

run-grpc:
	go run cmd/grpc/main.go

run-graphql:
	go run cmd/graphql/main.go

run-cmd:
	go run cmd/main.go

format:
	go fmt ./...

rest-dev:
	DC_APP_ENV=dev $(GOPATH)/bin/reflex -s -r '\.go$$' make format run-rest

grpc-dev:
	DC_APP_ENV=dev $(GOPATH)/bin/reflex -s -r '\.go$$' make format run-grpc

graphql-dev:
	DC_APP_ENV=dev $(GOPATH)/bin/reflex -s -r '\.go$$' make format run-graphql

test-cov:
	go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html

generate-swag:
	swag init -g main.go

generate-jwt-secret:
	$(eval JWT_SECRET := $(shell openssl rand -base64 32))
	@echo "$(JWT_SECRET)"

migration-up:
	migrate -database "mysql://root:root@tcp(localhost:3306)/clean-architecture" -path migrations up

migration-down:
	migrate -database "mysql://root:root@tcp(localhost:3306)/clean-architecture" -path migrations down

migration $$(enter):
	@read -p "Migration name:" migration_name; \
	migrate create -ext sql -dir migrations $$migration_name

proto:
	protoc -I ./pkg/protobuf --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. ./pkg/protobuf/*.proto