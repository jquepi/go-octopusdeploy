package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type SmtpConfigurationService struct {
	isConfiguredPath string

	services.Service
}

func NewSmtpConfigurationService(sling *sling.Sling, uriTemplate string, isConfiguredPath string) *SmtpConfigurationService {
	return &SmtpConfigurationService{
		isConfiguredPath: isConfiguredPath,
		Service:          services.NewService(constants.ServiceSMTPConfigurationService, sling, uriTemplate),
	}
}
