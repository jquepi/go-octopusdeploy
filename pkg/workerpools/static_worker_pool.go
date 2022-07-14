package workerpools

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type StaticWorkerPool struct {
	workerPool
}

type StaticWorkerPools struct {
	Items []*StaticWorkerPool `json:"Items"`
	resources.PagedResults
}

// NewStaticWorkerPool creates and initializes a static worker pool.
func NewStaticWorkerPool(name string) *StaticWorkerPool {
	return &StaticWorkerPool{
		workerPool: *newWorkerPool(name, WorkerPoolTypeStatic),
	}
}

// Validate checks the state of the static worker pool and returns an error if
// invalid.
func (s *StaticWorkerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(s)
}
