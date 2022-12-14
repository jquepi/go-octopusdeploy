package packages

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type PackageMetadataService struct {
	services.Service
}

func NewPackageMetadataService(sling *sling.Sling, uriTemplate string) *PackageMetadataService {
	return &PackageMetadataService{
		Service: services.NewService(constants.ServicePackageMetadataService, sling, uriTemplate),
	}
}
