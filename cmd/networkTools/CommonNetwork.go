package networkTools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"log"
	"strings"
)

func GetNetworkByName(client *gophercloud.ServiceClient, name string) []networks.Network {
	listOpts := networks.ListOpts{
		Name: name,
	}

	allPages, err := networks.List(client, listOpts).AllPages()
	if err != nil {
		log.Panic(err)
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		log.Panic(err)
	}

	return allNetworks
}

func GetTenantOpsNet(client *gophercloud.ServiceClient) networks.Network {
	var n networks.Network
	networks := GetNetworkByName(client, "")
	for _, network := range networks {
		if strings.Contains(network.Name, "ops-net") {
			n = network
		}
	}
	return n
}