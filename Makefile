GOPATH ?= $(HOME)/go

run-api:
	go run main.go

run-cmd:
	go run cmd/main.go

format:
	go fmt ./...

dev:
	DC_APP_ENV=dev $(GOPATH)/bin/reflex -s -r '\.go$$' make format run-api

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