package documents

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/compiler"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/preprocessor"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository/filesystem"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services"
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build document specified by ",
	Long:  ``,
	// TODO: Support building a list of lib
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cn, err := lib.ParseControlNumber(args[0])
		if err != nil {
			log.Fatal().Stack().Err(err).Msgf("%s is not a valid control number", args[0])
		}

		builder := services.NewBuilder(
			viper.Get("repo").(filesystem.FileRepository),
			preprocessor.NewPandocPreprocessor("markdown", "latex"),
			&template.LatexTemplate{},
			compiler.NewPdfLatexCompiler(viper.GetString("texinputs")),
		)

		err = builder.Build(cn)
		if err != nil {
			log.Fatal().Stack().Err(err).Msgf("Failed to build %s", args[0])
		}
	},
}

func init() {
	DocsCmd.AddCommand(buildCmd)
}
