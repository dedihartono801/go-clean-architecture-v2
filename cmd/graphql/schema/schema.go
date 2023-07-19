package schema

import (
	graphqlResolver "github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/graphql"
	"github.com/graphql-go/graphql"
)

// GenerateSchema will create a GraphQL Schema and set the Resolvers found in the UserService
// For all the needed fields
func GenerateSchema(ur graphqlResolver.UserResovler, jr graphqlResolver.JobResovler) (*graphql.Schema, error) {
	userType := generateUserType(jr)
	users := createUsersFields(userType, ur)
	user := createUserFields(userType, ur)
	// RootQuery
	fields := graphql.Fields{
		// We define the user query
		"users": &users,
		"user":  &user,
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	// RootMutation
	rootMutation := GenerateRootMutation(jr)

	// Now combine all Objects into a Schema Configuration
	schemaConfig := graphql.SchemaConfig{
		// Query is the root object query schema
		Query: graphql.NewObject(rootQuery),
		// Appliy the Mutation to the schema
		Mutation: rootMutation,
	}
	// Create a new GraphQL Schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
