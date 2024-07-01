package goraxauth

import (
	"os"

	"github.com/gophercloud/gophercloud/v2"
	tokens2 "github.com/gophercloud/gophercloud/v2/openstack/identity/v2/tokens"
)

var nilOptions = AuthOptions{}

/*
AuthOptionsFromEnv fills out a goraxauth.AuthOptionsRax structure with the
settings found on the various OpenStack OS_* environment variables and/or
the RAX_API_KEY environment variable.

The following variables provide sources of truth: OS_AUTH_URL, OS_USERNAME,
OS_PASSWORD, RAX_API_KEY and OS_PROJECT_ID.

Of these, OS_USERNAME and OS_PASSWORD or RAX_API_KEY must have settings,
or an error will result, while OS_PROJECT_ID is optional.

OS_TENANT_ID is the deprecated forms of OS_PROJECT_ID.

If OS_PROJECT_ID, they will still be referred as "tenant".

To use this function, first set the OS_* environment variables (for example,
by sourcing an `openrc` file), then:

	opts, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(context.TODO(), opts)
*/
func AuthOptionsFromEnv() (AuthOptions, error) {
	authURL := "https://identity.api.rackspacecloud.com/v2.0/"
	username := os.Getenv("OS_USERNAME")
	password := os.Getenv("OS_PASSWORD")
	raxApiKey := os.Getenv("RAX_API_KEY")
	tenantID := os.Getenv("OS_TENANT_ID")

	// if OS_AUTH_URL is set, override the authURL value.
	if v := os.Getenv("OS_AUTH_URL"); v != "" {
		authURL = v
	}

	// If OS_PROJECT_ID is set, overwrite tenantID with the value.
	if v := os.Getenv("OS_PROJECT_ID"); v != "" {
		tenantID = v
	}

	if username == "" {
		err := gophercloud.ErrMissingEnvironmentVariable{
			EnvironmentVariable: "OS_USERNAME",
		}
		return nilOptions, err
	}

	if raxApiKey == "" && password == "" {
		err := gophercloud.ErrMissingAnyoneOfEnvironmentVariables{
			EnvironmentVariables: []string{"OS_PASSWORD", "RAX_API_KEY"},
		}
		return nilOptions, err
	}

	ao := AuthOptions{
		AuthOptions: tokens2.AuthOptions{
			IdentityEndpoint: authURL,
			Username:         username,
			Password:         password,
			TenantID:         tenantID,
		},
		ApiKey: raxApiKey,
	}

	return ao, nil
}
