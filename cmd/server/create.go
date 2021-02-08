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
	"github.com/hasanaburayyan/openstack-tools/cmd/networkTools"
	ServerTools "github.com/hasanaburayyan/openstack-tools/cmd/serverTools"
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
		imageName, _ := cmd.Flags().GetString("image")
		flavorName, _ := cmd.Flags().GetString("flavor")
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

		server := ServerTools.CreateServer(serverClient, serverName, imageName, flavorName, networks)

		if volumeName != "" {
			volume := ServerTools.CreateVolume(volumeClient, volumeName)
			ServerTools.AttachVolume(serverClient, volume, server)
		}
	},
}

func init() {

	createCmd.PersistentFlags().StringVarP(&image, "image", "i", "", "Image to use")
	createCmd.PersistentFlags().StringVarP(&flavor, "flavor", "f", "", "Flavor to use")
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
