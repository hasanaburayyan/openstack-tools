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
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	ServerTools "github.com/hasanaburayyan/my-openstack-tools/cmd/serverTools"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverID, _ := cmd.Flags().GetString("id")
		fmt.Printf("deleting server %s\n", serverID)
		serverClient := ServerTools.GetOSClient()
		volumeClient := ServerTools.GetBlockStorageClient()
		server, _ := servers.Get(serverClient, serverID).Extract()
		ServerTools.DeleteServer(serverClient, serverID)
		deleteVolume, _ := cmd.Flags().GetBool("delete-volume")
		if deleteVolume {
			deleteOpts := volumes.DeleteOpts{
				Cascade: true,
			}
			volumes.Delete(volumeClient, server.AttachedVolumes[0].ID, deleteOpts)
			volumes.WaitForStatus(volumeClient, server.AttachedVolumes[0].ID, "", 600)
			fmt.Println("Volume was deleted")
		}
	},
}

var id string
var deleteVolume bool

func init() {
	ServerCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringVar(&id, "id", "", "The ID of the server you want to delete")
	deleteCmd.Flags().BoolP("delete-volume", "d", false, "Passed if you wish to delete the volume attached to the server")
	deleteCmd.MarkFlagRequired("id")
}
