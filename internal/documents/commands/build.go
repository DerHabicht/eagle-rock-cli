package commands

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/orchestrations"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a document",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("A control number must be provided to build a document.")
			os.Exit(1)
		}

		var orch orchestrations.BuildOrchestration
		var err error
		switch args[0][:strings.Index(args[0], "-")] {
		case "mr":
			orch, err = orchestrations.NewMemoForRecordBuildOrchestration()
			if err != nil {
				log.Fatal().Stack().Err(err).Msgf("Failed to initialize builder for %s", args[0])
			}
		case "warno":
			orch, err = orchestrations.NewWarnoBuildOrchestration()
			if err != nil {
				log.Fatal().Stack().Err(err).Msgf("Failed to initialize builder for %s", args[0])
			}
		default:
			log.Fatal().Msgf("Unsupported control number: %s", args[0])
		}

		err = orch.Build(args[0])
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("")
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
