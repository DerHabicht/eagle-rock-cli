package commands

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/orchestrations"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build one or more documents",
	Long:  ``,
	/*
		Run: func(cmd *cobra.Command, args []string) {
		},
	*/
}

var buildMemoForRecordCmd = &cobra.Command{
	Use:   "mr",
	Short: "Build a Memorandum for Record",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("A control number must be provided to build an MR.")
			os.Exit(1)
		}

		contentDir, ok := viper.Get("content").(map[string]interface{})["mr_dir"].(string)
		if !ok {
			log.Fatal().Msgf("Configuration key `content` is malformed: %#v", viper.Get("content"))
		}

		publishDir, ok := viper.Get("published").(map[string]interface{})["mr_dir"].(string)
		if !ok {
			log.Fatal().Msgf("Configuration key `published` is malformed: %#v", viper.Get("published"))
		}

		templateDir, ok := viper.Get("latex_templates").(map[string]interface{})["mr"].(string)
		if !ok {
			log.Fatal().Msgf("Configuration key `latex_templates` is malformed: %#v", viper.Get("latex_templates"))
		}

		texTemplate, err := ioutil.ReadFile(templateDir)
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("Failed to read LaTeX template for Memoranda for Record.")
		}

		orch := orchestrations.NewBuildOrchestration(
			services.NewFileContentReader(contentDir),
			services.NewMemoForRecordHeaderParser(),
			services.NewNullMacroProcessor(),
			services.NewPandocPreprocessor("latex"),
			services.NewLatexTemplate(texTemplate),
			services.NewPdflatexBuilder(publishDir),
		)

		err = orch.Build(args[0])
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("")
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// TODO: Switch build based on control number instead of specific command
	buildCmd.AddCommand(buildMemoForRecordCmd)
}
