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
	ServerTools "github.com/hasanaburayyan/openstack-tools/cmd/serverTools"
	"log"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details about server",
	Long: `
show server details
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Panic(err)
		}

		exactMatch, err := cmd.Flags().GetBool("exact-match")
		if err != nil {
			log.Panic(err)
		}

		serviceClient := ServerTools.GetOSClient()
		fmt.Printf("Searching for details on server %s...\n", name)

		if exactMatch {
			server, err := ServerTools.FindServerByExactName(serviceClient, name)
			if err != nil {
				log.Panic(err)
			}
			fmt.Printf("Found matching server with id: %s\n", server.ID)
			fmt.Println(server)
		} else {
			servers := ServerTools.FindServersByName(serviceClient, name)

			fmt.Printf("Found %d servers matching that name!..\n", len(servers))
			for _, server := range servers {
				fmt.Println(server.Name)
			}
		}
	},
}

func init() {
	ServerCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	showCmd.Flags().StringVarP(&name, "name", "n", "", "Namae of server to show details about")
	showCmd.Flags().BoolP("exact-match", "e", false, "Toggle this flag to search for only a single server with an exact name match")
}
