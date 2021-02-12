/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package server

import (
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/hasanaburayyan/my-openstack-tools/cmd/flavorTools"
	"github.com/hasanaburayyan/my-openstack-tools/cmd/imageTools"
	"github.com/hasanaburayyan/my-openstack-tools/cmd/networkTools"
	ServerTools "github.com/hasanaburayyan/my-openstack-tools/cmd/serverTools"
	"github.com/spf13/cobra"
	"log"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Long: `
The create command allows you to create new VMs with required and optional flags
see usage below
`,
	Run: func(cmd *cobra.Command, args []string) {
		serverName, _ := cmd.Flags().GetString("name")
		imageID, _ := cmd.Flags().GetString("imageID")
		imageName, _ := cmd.Flags().GetString("imageName")
		flavorID, _ := cmd.Flags().GetString("flavorID")
		flavorName, _ := cmd.Flags().GetString("flavorName")
		volumeName, _ := cmd.Flags().GetString("volumeName")
		networkName, _ := cmd.Flags().GetString("networkName")

		serverClient := ServerTools.GetOSClient()
		volumeClient := ServerTools.GetBlockStorageClient()
		networkClient := networkTools.GetNetworkClient()

		// Check If Server Already Exists
		s, _ := ServerTools.FindServerByExactName(serverClient, serverName)
		if s.ID != "" {
			log.Panic(fmt.Sprintf("Server(s) Matching With Name: %s Already Exist!!", serverName))
		}

		// Retrieve Networks To Add To Server
		networks := networkTools.GetNetworkByName(networkClient, networkName)
		if len(networks) == 0 {
			log.Panic(fmt.Sprintf("Cannot Find A Network Matching Name: %s!\n", networkName))
		}

		createOpts := servers.CreateOpts{}
		createOpts.Name = serverName
		ServerTools.AttachNetworkToOpts(&createOpts, networks)
		if imageID != "" {
			createOpts.ImageRef = imageID
		} else if imageName != "" {
			image, err := imageTools.GetImageByName(serverClient, imageName)
			if err != nil {
				log.Panic(err)
			}
			createOpts.ImageRef = image.ID
		}

		if flavorID != "" {
			createOpts.FlavorRef = flavorID
		} else if flavorName != "" {
			flavor, err := flavorTools.GetFlavorByName(serverClient, flavorName)
			if err != nil {
				log.Panic(err)
			}
			createOpts.FlavorRef = flavor.ID
		}

		//server := ServerTools.CreateServer(serverClient, serverName, imageID, flavorID, networks)
		server := ServerTools.CreateServerWithOptions(serverClient, createOpts)

		if volumeName != "" {
			volume := ServerTools.CreateVolume(volumeClient, volumeName)
			ServerTools.AttachVolume(serverClient, volume, server)
		}
	},
}

func init() {

	createCmd.PersistentFlags().StringVarP(&imageID, "imageID", "i", "", "Image ID to use (Cannot Be Used With ImageName")
	createCmd.PersistentFlags().StringVar(&imageName, "imageName", "", "Image Name To Use (Cannot Be Used With ImageID")
	createCmd.PersistentFlags().StringVarP(&flavorID, "flavorID", "f", "", "Flavor ID to use (Cannot Be Used With FlavorName")
	createCmd.PersistentFlags().StringVar(&flavorName, "flavorName", "", "Flavor Name To Use (Cannot Be Used With FlavorID")
	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name Of Server")
	createCmd.PersistentFlags().StringVarP(&volumeName, "volumeName", "v", "", "Name of volume")
	createCmd.PersistentFlags().StringVarP(&netName, "networkName", "x", "", "Name of Network")
	createCmd.MarkPersistentFlagRequired("image")
	createCmd.MarkPersistentFlagRequired("flavor")
	createCmd.MarkPersistentFlagRequired("name")
	createCmd.MarkPersistentFlagRequired("networkName")

	ServerCmd.AddCommand(createCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
