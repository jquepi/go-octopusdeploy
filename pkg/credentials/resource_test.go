package credentials

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestResourceWithUsernamePasswordAsJSON(t *testing.T) {
	description := internal.GetRandomName()
	id := internal.GetRandomName()
	name := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())
	selfLink := internal.GetRandomName()
	username := internal.GetRandomName()

	usernamePassword := NewUsernamePassword(username, password)

	usernamePasswordAsJSON, err := json.Marshal(usernamePassword)
	require.NoError(t, err)
	require.NotNil(t, usernamePasswordAsJSON)

	resource := NewResource(name, usernamePassword)
	resource.Description = description
	resource.ID = id
	resource.Links["Self"] = selfLink

	expectedJSON := fmt.Sprintf(`{
		"Description": "%s",
		"Details": %s,
		"Id": "%s",
		"Name": "%s",
		"Links": {
			"Self": "%s"
		}
	}`, description, usernamePasswordAsJSON, id, name, selfLink)

	resourceAsJSON, err := json.Marshal(resource)
	require.NoError(t, err)
	require.NotNil(t, resourceAsJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(resourceAsJSON))
}

func TestResourceWithAnonymousAsJSON(t *testing.T) {
	anonymous := NewAnonymous()
	description := internal.GetRandomName()
	id := internal.GetRandomName()
	name := internal.GetRandomName()
	selfLink := internal.GetRandomName()

	anonymousdAsJSON, err := json.Marshal(anonymous)
	require.NoError(t, err)
	require.NotNil(t, anonymousdAsJSON)

	resource := NewResource(name, anonymous)
	resource.Description = description
	resource.ID = id
	resource.Links["Self"] = selfLink

	expectedJSON := fmt.Sprintf(`{
		"Description": "%s",
		"Details": %s,
		"Id": "%s",
		"Name": "%s",
		"Links": {
			"Self": "%s"
		}
	}`, description, anonymousdAsJSON, id, name, selfLink)

	resourceAsJSON, err := json.Marshal(resource)
	require.NoError(t, err)
	require.NotNil(t, resourceAsJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(resourceAsJSON))
}
