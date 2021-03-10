package documents

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository/filesystem"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Build document index",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		indexer := services.NewIndexer(viper.Get("repo").(filesystem.FileRepository))

		list, err := indexer.IndexDocuments()
	},
}

func init() {
	DocsCmd.AddCommand(indexCmd)
}
