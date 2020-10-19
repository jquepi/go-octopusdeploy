package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type rootService struct {
	service
}

func newRootService(sling *sling.Sling, uriTemplate string) *rootService {
	return &rootService{
		service: newService(serviceRootService, sling, uriTemplate, new(model.RootResource)),
	}
}

func (s rootService) Get() (*model.RootResource, error) {
	path, err := getPath(s)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.RootResource), nil
}

var _ IService = &rootService{}
