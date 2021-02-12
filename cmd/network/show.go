package network

import (
	"fmt"
	"github.com/hasanaburayyan/my-openstack-tools/cmd/networkTools"
	"github.com/spf13/cobra"
	"log"
)

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show Information About A Openstack Network",
	Long: `
Show Useful Information About An Openstack Network
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Panic(err)
		}

		client := networkTools.GetNetworkClient()

		n := networkTools.GetNetworkByName(client, name)

		for _, network := range n {
			fmt.Printf("Name: %s\tID: %s\n",network.Name, network.ID)
		}
	},
}

var name string

func init() {
	NetworkCmd.AddCommand(showCmd)

	// Local Flags
	showCmd.Flags().StringVarP(&name, "name", "n", "", "Name Of Network To Show")
}