package image

import (
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/hasanaburayyan/my-openstack-tools/cmd/imageTools"
	ServerTools "github.com/hasanaburayyan/my-openstack-tools/cmd/serverTools"
	"github.com/spf13/cobra"
	"log"
)

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show image details",
	Long: `
Shows image details
`,
	Run: func(cmd *cobra.Command, args []string) {
		serverClient := ServerTools.GetOSClient()
		imageName, _ := cmd.Flags().GetString("name")
		imageID, _ := cmd.Flags().GetString("id")

		var image images.Image
		var err error

		if imageID != "" {
			image = imageTools.GetImageByID(serverClient, imageID)
		} else if imageName != "" {
			image, err = imageTools.GetImageByName(serverClient, imageName)
			if err != nil {
				log.Panic(err)
			}
		}

		fmt.Printf("==> %s\t||\t%s", image.Name, image.ID)
	},
}

func init() {
	ImageCmd.AddCommand(showCmd)

	showCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the image to show details of")
	showCmd.Flags().StringVarP(&id, "id", "i", "", "The ID of the image to show details on")
}
