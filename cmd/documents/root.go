package documents

import (
	"github.com/spf13/cobra"
)

var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Perform an operation on one or more policy documents",
	Long:  ``,
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new memo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a policy memo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	DocsCmd.AddCommand(newCmd)
	DocsCmd.AddCommand(updateCmd)
}
