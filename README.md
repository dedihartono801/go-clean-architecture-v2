## Description

[Go clean architecture]

boilerplate modification from https://github.com/cmparrela/go-clean-architecture

Examples of types of communication;

- API
- CLI

Examples of data persistence;

- Mysql
- Mongo
- In memory
- redis

Example:

- rest api
- Dockerfile with multi stage build (for prod/staging)
- DockerfileDev (for local environment)
- Docker-compose
- swagger for doc
- middleware auth jwt
- migration
- unit testing with mock and table test
- concurrency with goroutines by implementing mutex and locking rows to avoid race conditions (on checkout api)
- worker queue using asynq library

## Install Swagger

go install github.com/swaggo/swag/cmd/swag@latest

## Install Migration

go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## Run Service

```bash
$ docker-compose up -d
```

## Run Migration UP

```bash
$ make migration-up
```

## Run Migration Down

```bash
$ make migration-down
```

## Create Migration

```bash
$ make migration
#type your migration name example: create_create_table_users
```

## Create .env file

```bash
$ ./entrypoint.sh
```

## Generate JWT Secret

```bash
#install openssl on your OS and run command below
$ make generate-jwt-secret
#copy the secret key and then create new env called JWT_SECRET in .env file:
```

## Generate Swagger

```bash
$ make generate-swag
```

## Test Coverage

```bash
$ make test-cov
```

## CLI Documentation

```bash
#entering go app Docker container
$ docker exec -it go-app /bin/sh
#create user
$ go run cmd/main.go user create -n=teste -e=teste@gmail.com
#update user
$ go run cmd/main.go user update -n=teste -e=teste@gmail.com -i=9cc26bf0-1272-45c8-93c5-1d83cfe82033
```

## API Documentation

API Documentation was created with swagger and is available at `http://localhost:5001/docs`

## Fiber Monitoring

Available at `http://localhost:5001`
