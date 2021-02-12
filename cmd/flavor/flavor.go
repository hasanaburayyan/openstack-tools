package flavor

import "github.com/spf13/cobra"

var FlavorCmd = &cobra.Command{
	Use: "flavor",
	Short: "Useful Commands For Interacting With Flavors In Openstack",
	Long: `
Useful Commands For Interacting With Flavors In Openstack
`,
}

var (
	name string
)

func init() {

}