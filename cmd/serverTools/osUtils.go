package ServerTools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"log"
)

func GetOSClient() *gophercloud.ServiceClient {
	opts, err := openstack.AuthOptionsFromEnv()

	if err != nil {
		log.Panic(err)
	}

	// Create a provider to authenticate all services we use
	provider, err := openstack.AuthenticatedClient(opts)

	if err != nil {
		log.Panic(err)
	}

	// Use the provider to authenticate a new client
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Type: "compute",
		Region: "RegionOne",
	})

	if err != nil {
		log.Panic(err)
	}
	return client
}
