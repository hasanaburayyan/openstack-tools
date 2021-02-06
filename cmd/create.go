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
package cmd

import (
	"fmt"
	ServerTools "github.com/hasanaburayyan/openstack-tools/cmd/serverTools"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverName, _ := cmd.Flags().GetString("name")
		imageName, _ := cmd.Flags().GetString("image")
		flavorName, _ := cmd.Flags().GetString("flavor")

		fmt.Printf("create called with \nname: %s\nimage: %s\nflavor: %s\n", serverName, imageName, flavorName)

		client := ServerTools.GetOSClient()
		ServerTools.CreateServer(client, serverName, imageName, flavorName)


	},
}





func init() {


	createCmd.PersistentFlags().StringVarP(&image, "image", "i", "", "Image to use")
	createCmd.PersistentFlags().StringVarP(&flavor, "flavor", "f", "", "Flavor to use")
	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name Of Server")

	createCmd.MarkPersistentFlagRequired("image")
	createCmd.MarkPersistentFlagRequired("flavor")
	createCmd.MarkPersistentFlagRequired("name")

	serverCmd.AddCommand(createCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
