package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// TODO: Add doc index commands
// TODO: Add build commands

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "documents",
	Short: "CLI tool to manage memoranda for record",
	Long: `CLI tool for managing memoranda for record.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
