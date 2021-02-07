package ServerTools

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func ListServersInCurrentTenant(client *gophercloud.ServiceClient) {
	// Options for listing servers
	listOpts := servers.ListOpts{
		AllTenants: false,
	}

	// Get all pages of servers
	allPages, err := servers.List(client, listOpts).AllPages()

	if err != nil {
		log.Panic(err)
	}

	// Extract all servers from pages
	allServers, err := servers.ExtractServers(allPages)

	if err != nil {
		log.Panic(err)
	}

	// Print out ever server by name
	for _, server := range allServers {
		fmt.Printf("Name: %s, Image: %s, Flavor: %s, Networks: %s\n", server.Name, server.Image["id"], server.Flavor["id"], server.Addresses)
	}
}

func ListAllKeypairs(client *gophercloud.ServiceClient) {
	allPages, err := keypairs.List(client).AllPages()
	if err != nil {
		log.Panic(err)
	}

	allKeyPairs, err := keypairs.ExtractKeyPairs(allPages)
	if err != nil {
		log.Panic(err)
	}

	for _, kp := range allKeyPairs {
		fmt.Println(kp)
	}
}

func DeleteServer(client *gophercloud.ServiceClient, serverId string) {
	err := servers.Delete(client, serverId).ExtractErr()
	if err != nil {
		panic(err)
	}
	fmt.Println("Waiting For server to delete will timeout in 600 seconds")
	servers.WaitForStatus(client, serverId, "", 600)
	fmt.Printf("server %s deleted!", serverId)
}

func CreateServer(client *gophercloud.ServiceClient, serverName, imageName, flavorName string) *servers.Server {
	serverCreateOpts := servers.CreateOpts{
		Name:      serverName,
		ImageRef:  imageName,
		FlavorRef: flavorName,
		Networks: []servers.Network{
			servers.Network{UUID: "6353f4fd-0ec8-43cb-aedd-d575d8db1721"},
		},
		Metadata: map[string]string{
			"hsa29-test": "true",
		},
	}

	createOpts := keypairs.CreateOptsExt{
		CreateOptsBuilder: serverCreateOpts,
		KeyName:           "opskey",
	}

	server, err := servers.Create(client, createOpts).Extract()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Waiting For server to transition to ACTIVE, will timeout in 600 seconds")
	servers.WaitForStatus(client, server.ID, "ACTIVE", 600)
	fmt.Printf("server %s created!\n", server.ID)

	return server
}

func CreateVolume(client *gophercloud.ServiceClient, volumeName string) *volumes.Volume {
	opts := volumes.CreateOpts{Size: 10, Name: volumeName, VolumeType: "d4559dc6-3abc-49a1-aed6-a2c4f0b4ceac"}
	vol, err := volumes.Create(client, opts).Extract()

	if err != nil {
		log.Panic(err)
	}
	volumes.WaitForStatus(client, vol.ID, "available", 600)
	fmt.Printf("volume %s created!\n", volumeName)
	return vol
}

func AttachVolume(client *gophercloud.ServiceClient, volume *volumes.Volume, server *servers.Server) {
	fmt.Printf("Attempting to attach volume %s to %s!\n", volume.ID, server.ID)
	createOpts := volumeattach.CreateOpts{
		Device:   "/dev/vdc",
		VolumeID: volume.ID,
	}

	_, err := volumeattach.Create(client, server.ID, createOpts).Extract()
	if err != nil {
		panic(err)
	}
	volumes.WaitForStatus(client, volume.ID, "in-use", 600)
}
