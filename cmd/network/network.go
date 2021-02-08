package network

import "github.com/spf13/cobra"

var NetworkCmd = &cobra.Command{
	Use: "network",
	Short: "Useful Openstack Network Commands",
	Long:`
A series of Useful Network Commands for Openstack
`,
}

func init() {

}