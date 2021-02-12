package flavor

import (
	"fmt"
	"github.com/hasanaburayyan/my-openstack-tools/cmd/flavorTools"
	ServerTools "github.com/hasanaburayyan/my-openstack-tools/cmd/serverTools"
	"github.com/spf13/cobra"
	"log"
)

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show Details About Flavor",
	Long: `
Show details on openstack flavors
`,
	Run: func(cmd *cobra.Command, args []string) {
		serverClient := ServerTools.GetOSClient()
		flavorName, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Searching for flavor with name: %s\n", flavorName)

		image, err := flavorTools.GetFlavorByName(serverClient, flavorName)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Name: %s\nID: %s\nDisk: %d\nRAM: %d\nVCPU: %d\n", image.Name, image.ID, image.Disk, image.RAM, image.VCPUs)
	},
}


func init() {
	FlavorCmd.AddCommand(showCmd)

	showCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the flavor to show")
}