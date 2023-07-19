package job

import (
	"errors"
	"fmt"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	"github.com/graphql-go/graphql"
)

type Service interface {
	GetJobs(p graphql.ResolveParams) ([]entity.Job, error)
	MutateJobs(p graphql.ResolveParams) (*entity.Job, error)
}

type service struct {
	repository repository.MemoryJobRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewGraphqlJobService(
	repository repository.MemoryJobRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) GetJobs(p graphql.ResolveParams) ([]entity.Job, error) {
	// Fetch Source Value
	g, ok := p.Source.(entity.User)

	if !ok {
		return nil, errors.New("source was not a Gopher")
	}
	// Here we extract the Argument Company
	company := ""
	if value, ok := p.Args["company"]; ok {
		company, ok = value.(string)
		if !ok {
			return nil, errors.New("id has to be a string")
		}
	}
	return s.repository.GetJobs(g.ID, company)
}

// MutateJobs is used to modify jobs based on a mutation request
// Available params are
// employeeid! -- the id of the employee, required
// jobid! -- job to modify, required
// start -- the date to set as start date
// end -- the date to set as end
func (gs *service) MutateJobs(p graphql.ResolveParams) (*entity.Job, error) {
	employee, err := grabStringArgument("employeeid", p.Args, true)
	if err != nil {
		return nil, err
	}
	jobid, err := grabStringArgument("jobid", p.Args, true)
	if err != nil {
		return nil, err
	}
	start, err := grabStringArgument("start", p.Args, false)
	if err != nil {
		return nil, err
	}
	end, err := grabStringArgument("end", p.Args, false)
	if err != nil {
		return nil, err
	}

	// Get the job
	job, err := gs.repository.GetJob(employee, jobid)
	if err != nil {
		return nil, err
	}
	// Modify start and end date if they are set
	if start != "" {
		job.Start = start
	}

	if end != "" {
		job.End = end
	}
	// Update with new values
	job, err = gs.repository.Update(job)
	if err != nil {
		return nil, err
	}
	return &job, nil

}

// grabStringArgument is used to grab a string argument
func grabStringArgument(k string, args map[string]interface{}, required bool) (string, error) {
	// first check presense of arg
	if value, ok := args[k]; ok {
		// check string datatype
		v, o := value.(string)
		if !o {
			return "", fmt.Errorf("%s is not a string value", k)
		}
		return v, nil
	}
	if required {
		return "", fmt.Errorf("missing argument %s", k)
	}
	return "", nil
}
