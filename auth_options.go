package goraxauth

import (
	"context"
	"fmt"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	tokens2 "github.com/gophercloud/gophercloud/v2/openstack/identity/v2/tokens"
)

type AuthOptions struct {
	tokens2.AuthOptions
	ApiKey string `json:"apiKey,omitempty"`
}

func (opts AuthOptions) ToTokenV2CreateMap() (map[string]interface{}, error) {
	// if we have an ApiKey, use that otherwise just use the regular auth mechanism
	if opts.ApiKey != "" {
		if opts.Username == "" {
			return nil, fmt.Errorf("username must be supplied for API key auth")
		}
		return map[string]interface{}{
			"auth": map[string]interface{}{
				"RAX-KSKEY:apiKeyCredentials": map[string]interface{}{
					"username": opts.Username,
					"apiKey":   opts.ApiKey,
				},
			},
		}, nil
	} else if (opts.Username != "" && opts.Password != "") || opts.TokenID != "" {
		return opts.AuthOptions.ToTokenV2CreateMap()
	} else {
		return nil, fmt.Errorf("missing (Username and (Password or ApiKey)) or TokenId for auth")
	}
}

func (opts AuthOptions) CanReauth() bool {
	return opts.AllowReauth
}

func AuthenticatedClient(ctx context.Context, options AuthOptions) (*gophercloud.ProviderClient, error) {
	client, err := openstack.NewClient(options.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	err = openstack.AuthenticateV2(ctx, client, options, gophercloud.EndpointOpts{})
	if err != nil {
		return nil, err
	}
	return client, nil
}
