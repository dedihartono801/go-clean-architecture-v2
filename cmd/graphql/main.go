// This package is a demonstration how to build and use a GraphQL server in Go
package main

import (
	"log"
	"net/http"

	"github.com/dedihartono801/go-clean-architecture-v2/cmd/graphql/schema"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/graphql/job"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/graphql/user"
	graphqlHandler "github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/graphql"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	//envConfig := config.SetupEnvFile()
	//mysql := database.InitMysql(envConfig)

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	jobRepository := repository.NewMemoryJobRepository()
	userRepository := repository.NewMemoryUserRepository()
	userService := user.NewGraphqlUserService(userRepository, validator, identifier)
	jobService := job.NewGraphqlJobService(jobRepository, validator, identifier)
	userResolver := graphqlHandler.NewUserResolver(userService)
	jobResolver := graphqlHandler.NewJobResolver(jobService)

	schema, err := schema.GenerateSchema(userResolver, jobResolver)
	if err != nil {
		panic(err)
	}

	StartServer(schema)
}

// StartServer will trigger the server with a Playground
func StartServer(schema *graphql.Schema) {
	// Create a new HTTP handler
	h := handler.New(&handler.Config{
		Schema: schema,
		// Pretty print JSON response
		Pretty: true,
		// Host a GraphiQL Playground to use for testing Queries
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":5003", nil))
}
