# goraxauth

A [gophercloud][gophercloud] compatible authentication helper
for [Rackspace Customer Identity][raxident].

This initial version only adds support for the Rackspace API KEY and
does not support password authentication for MFA enabled accounts.

## Usage

Replace any usage of
[github.com/gophercloud/gophercloud/openstack/identity/v2/tokens.AuthOptions][tokens2-authoptions]
with `goraxauth.AuthOptions`. This struct adds the `ApiKey` field which
is the users Rackspace API Key.

The following functions from the [docs for gophercloud on go.dev][go-gophercloud]
examples are replaced by functions from `goraxauth`.

```go
import (
    "github.com/gophercloud/gophercloud/v2/openstack"
    "github.com/rackerlabs/goraxauth"
)

// old functions
opts, err := openstack.AuthOptionsFromEnv()
provider, err := openstack.AuthenticatedClient(opts)

// replacements
opts, err := goraxauth.AuthOptionsFromEnv()
provider, err := goraxauth.AuthenticatedClient(opts)
```

[go-gophercloud]: <https://pkg.go.dev/github.com/gophercloud/gophercloud>
[gophercloud]: <https://github.com/gophercloud/gophercloud>
[raxident]: <https://docs.rackspace.com/docs/cloud-identity-v2-getting-started>
[tokens2-authoptions]: <https://pkg.go.dev/github.com/gophercloud/gophercloud/v2@v2.0.0/openstack/identity/v2/tokens#AuthOptions>
