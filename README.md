## Description

[Go clean architecture]

boilerplate modification from https://github.com/dedihartono801/go-clean-architecture

1. `cmd/`: This folder contains the application's entry point(s) or executable(s).

2. `internal/`: This folder holds the core application code. It is not accessible from outside the module/package.

   - `app/`: This folder contains the application-specific logic.

     - `usecase/`: Contains the application's use cases or business logic. For example, `user.go` could define use cases related to user management.

     - `repository/`: This folder contains interfaces or contracts that define how to interact with external dependencies, such as databases or APIs. For example, `user_repository.go` could define the methods for fetching and saving user data.

   - `entity/`: This folder defines the application's entities or domain models. For example, `user.go` could define the user entity with its properties and behavior.

   - `delivery/`: Contains the delivery mechanisms, such as HTTP handlers, used to interact with the outside world.

     - `http/`: This folder contains the HTTP-specific code. For example, `user_handler.go` could define the HTTP handlers for user-related endpoints.

3. `pkg/`: This folder contains shared packages or utilities that can be used by different parts of the application. For example, `logger/logger.go` could define a logger package used throughout the application.

   - `config/`: This folder holds configuration-related code. For example, `config.go` could define functions or methods for loading application configuration.

4. `migrations/`: This folder may contain database migration scripts or related files.

5. `database/`: This folder may contain database-specific code or configurations.

Examples of types of communication:

- API
- CLI

Examples of data persistence:

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
