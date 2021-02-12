package commands

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Print and edit the document index.",
	Long:  ``,
}

var unpublishedIndexCmd = &cobra.Command{
	Use: "unpublished",
	Short: "Print the document index.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := models.LoadDocumentIndex(viper.GetString("index"))
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("Failed to load document index.")
		}

		var unpublished []string
		for k, v := range idx.Mr {
			for _, w := range v {
				if w.Date == nil {
					unpublished = append(unpublished, fmt.Sprintf("%s: %s - %s", k, w.ControlNumber, w.Subject))
				}
			}
		}

		for _, v := range unpublished {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.AddCommand(unpublishedIndexCmd)
}
