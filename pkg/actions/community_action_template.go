package actions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// CommunityActionTemplates defines a collection of community action templates
// with built-in support for paged results.
type CommunityActionTemplates struct {
	Items []*CommunityActionTemplate `json:"Items"`
	resources.PagedResults
}

// CommunityActionTemplate represents a community action template in Octopus
// Deploy.
type CommunityActionTemplate struct {
	ActionType  string                                    `json:"ActionType" validate:"required,notblank"`
	Author      string                                    `json:"Author,omitempty"`
	Category    string                                    `json:"Category,omitempty"`
	Description string                                    `json:"Description,omitempty"`
	ExternalId  *uuid.UUID                                `json:"ExternalId,omitempty"`
	HistoryURL  string                                    `json:"HistoryUrl,omitempty"`
	Name        string                                    `json:"Name" validate:"required"`
	Packages    []packages.PackageReference               `json:"Packages,omitempty"`
	Parameters  []actiontemplates.ActionTemplateParameter `json:"Parameters"`
	Properties  map[string]core.PropertyValue             `json:"Properties,omitempty"`
	Type        string                                    `json:"Type,omitempty"`
	Version     int32                                     `json:"Version,omitempty"`
	Website     string                                    `json:"Website,omitempty"`

	resources.Resource
}

// NewCommunityActionTemplate initializes a community action template.
func NewCommunityActionTemplate(name string, actionType string) *CommunityActionTemplate {
	return &CommunityActionTemplate{
		ActionType: actionType,
		Name:       name,
		Packages:   []packages.PackageReference{},
		Parameters: []actiontemplates.ActionTemplateParameter{},
		Properties: map[string]core.PropertyValue{},
		Resource:   *resources.NewResource(),
	}
}

// Validate checks the state of this CommunityActionTemplate and returns an
// error if invalid.
func (c *CommunityActionTemplate) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(c)
}
