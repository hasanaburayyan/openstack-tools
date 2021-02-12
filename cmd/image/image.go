package image

import "github.com/spf13/cobra"

var ImageCmd = &cobra.Command{
	Use: "image",
	Short: "Commands to interact with openstack images",
	Long: `
The following commands are useful tools for interacting with openstack images 
`,

}

var (
	name string
	id string
)

func init() {

}
