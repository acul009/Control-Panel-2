package deploymentsapi

import (
	"context"
	"fmt"
	"log"

	"github.com/acul009/control-mono/api/deployments/database"
	dbErrors "github.com/acul009/control-mono/api/deployments/database/errors"
	deployments "github.com/acul009/control-mono/api/deployments/gen/deployments"
)

// deployments service example implementation.
// The example methods log the requests and return zero values.
type deploymentssrvc struct {
	logger *log.Logger
}

// NewDeployments returns the deployments service implementation.
func NewDeployments(logger *log.Logger) deployments.Service {
	return &deploymentssrvc{logger}
}

var db database.Database = database.GetDatabase()

// Upsert implements upsert.
func (s *deploymentssrvc) Upsert(ctx context.Context, p *deployments.Deployment) (err error) {
	s.logger.Print("deployments.upsert")
	db.Save(p)
	return
}

// List implements list.
func (s *deploymentssrvc) List(ctx context.Context) (res []string, err error) {
	s.logger.Print("deployments.list")
	res = db.List()
	return
}

// Get implements get.
func (s *deploymentssrvc) Get(ctx context.Context, p string) (res *deployments.Deployment, err error) {
	var dbError error
	res, dbError = db.Load(p)
	if dbError != nil {
		switch dbError.(type) {
		case dbErrors.ErrNotFound:
			err = deployments.MakeDeploymentNotFound(dbError)
			break
		default:
			err = fmt.Errorf("Internal Server Error")
		}
	}
	return
}

// Delete implements delete.
func (s *deploymentssrvc) Delete(ctx context.Context, p string) (err error) {
	dbError := db.Delete(p)
	if dbError != nil {
		switch dbError.(type) {
		case dbErrors.ErrNotFound:
			err = deployments.MakeDeploymentNotFound(dbError)
			break
		default:
			err = fmt.Errorf("Internal Server Error")
		}
	}
	return
}
