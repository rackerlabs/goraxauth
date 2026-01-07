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

func (opts AuthOptions) ToTokenV2CreateMap() (map[string]any, error) {
	if opts.ApiKey != "" {
		if opts.Username == "" {
			return nil, fmt.Errorf("username is required when using API key authentication")
		}
		return map[string]any{
			"auth": map[string]any{
				"RAX-KSKEY:apiKeyCredentials": map[string]any{
					"username": opts.Username,
					"apiKey":   opts.ApiKey,
				},
			},
		}, nil
	}

	// Fall back to embedded or standard AuthOptions if no API key is provided.
	if (opts.Username != "" && opts.Password != "") || opts.TokenID != "" {
		return opts.AuthOptions.ToTokenV2CreateMap()
	}

	return nil, fmt.Errorf("authentication requires either: (Username and (Password or ApiKey)) or TokenID")
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
