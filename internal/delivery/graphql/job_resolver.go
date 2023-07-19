package graphql

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/graphql/job"
	"github.com/graphql-go/graphql"
)

type JobResovler interface {
	ResolveJobs(p graphql.ResolveParams) (interface{}, error)
	MutateJobs(p graphql.ResolveParams) (interface{}, error)
}

type jobResolver struct {
	service job.Service
}

func NewJobResolver(service job.Service) JobResovler {
	return &jobResolver{service}
}

// ResolveJobs is used to find all jobs related to a user
func (gs *jobResolver) ResolveJobs(p graphql.ResolveParams) (interface{}, error) {

	// Find Jobs Based on the Users ID
	jobs, err := gs.service.GetJobs(p)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// MutateJobs is used to modify jobs based on a mutation request
// Available params are
// employeeid! -- the id of the employee, required
// jobid! -- job to modify, required
// start -- the date to set as start date
// end -- the date to set as end
func (gs *jobResolver) MutateJobs(p graphql.ResolveParams) (interface{}, error) {
	job, err := gs.service.MutateJobs(p)
	if err != nil {
		return nil, err
	}
	return job, nil
}
