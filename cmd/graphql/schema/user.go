package schema

import (
	graphqlResolver "github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/graphql"

	"github.com/graphql-go/graphql"
)

// genereateUserType will assemble the Usertype and all related fields
func generateUserType(jr graphqlResolver.JobResovler) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		// Fields is the field values to declare the structure of the object
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "The ID that is used to identify unique user",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the user",
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "The email of the user",
			},
			// Here we create a graphql.Field which is depending on the jobs repository, notice how the User struct does not contain any information about jobs
			// But this still works
			"jobs": generateJobsField(jr),
		}})
}

// Fungsi untuk membuat field-field terkait User
func createUsersFields(userType graphql.Type, ur graphqlResolver.UserResovler) graphql.Field {
	return graphql.Field{
		Type:        graphql.NewList(userType),
		Resolve:     ur.ResolveUsers,
		Description: "Query all Users",
	}
}

// Fungsi untuk membuat field-field terkait user
func createUserFields(userType graphql.Type, ur graphqlResolver.UserResovler) graphql.Field {
	return graphql.Field{
		Type:        userType,
		Resolve:     ur.ResolveUser,
		Description: "Query Single user",
	}
}
