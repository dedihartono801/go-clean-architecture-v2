package schema

import (
	graphqlResolver "github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/graphql"
	"github.com/graphql-go/graphql"
)

// We can initialize Objects like this unless they need a special resolver
var JobType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Job",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"employeeID": &graphql.Field{
			Type: graphql.String,
		},
		"company": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"start": &graphql.Field{
			Type: graphql.String,
		},
		"end": &graphql.Field{
			Type: graphql.String,
		},
	},
},
)

// generateJobsField will build the GraphQL Field for jobs
func generateJobsField(gs graphqlResolver.JobResovler) *graphql.Field {
	return &graphql.Field{
		// Return a list of Jobs
		Type:        graphql.NewList(JobType),
		Description: "A list of all jobs the user had",
		Resolve:     gs.ResolveJobs,
		// Args are the possible arguments.
		Args: graphql.FieldConfigArgument{
			"company": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
	}
}

// modifyJobArgs are arguments available for the modifyJob Mutation request
var modifyJobArgs = graphql.FieldConfigArgument{
	"employeeid": &graphql.ArgumentConfig{
		// Create a string argument that cannot be empty
		Type: graphql.NewNonNull(graphql.String),
	},
	"jobid": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	// The new start date to apply if set
	"start": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	// The new end date to apply if set
	"end": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// generateRootMutation will create the root mutation object
func GenerateRootMutation(jr graphqlResolver.JobResovler) *graphql.Object {

	mutationFields := graphql.Fields{
		// Create a mutation named modifyJob which accepts a JobType
		"modifyJob": generateGraphQLField(JobType, jr.MutateJobs, "Modify a job for a user", modifyJobArgs),
	}
	mutationConfig := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}

	return graphql.NewObject(mutationConfig)
}

// generateGraphQLField is a generic builder factory to create graphql fields
func generateGraphQLField(output graphql.Output, resolver graphql.FieldResolveFn, description string, args graphql.FieldConfigArgument) *graphql.Field {
	return &graphql.Field{
		Type:        output,
		Resolve:     resolver,
		Description: description,
		Args:        args,
	}
}
