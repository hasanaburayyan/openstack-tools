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
	ServerTools "github.com/hasanaburayyan/my-openstack-tools/cmd/serverTools"
	"github.com/spf13/cobra"
	"strings"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all servers in tenant",
	Long: `
Lists all servers in tenant
`,
	Run: func(cmd *cobra.Command, args []string) {
		temp, _ := cmd.Flags().GetString("format")
		client := ServerTools.GetOSClient()
		ServerTools.ListServersInCurrentTenant(client, formatTemplate(temp))
	},
}

var template string

func formatTemplate(t string) string {
	t = strings.Replace(t, "\\t", "\t", -2)
	t = strings.Replace(t, "\\n", "\n", -2)
	return t
}

func init() {
	ServerCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringVarP(&template, "format", "f", "", "GO Template Format")

}
