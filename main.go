package main

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/cmd"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository/filesystem"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func setDefaultConfig(home string) {
	root := fmt.Sprintf("%s/Documents/Org/Policy", home)
	viper.SetDefault("repo_type", "filesystem")
	viper.SetDefault("src", map[string]string{
		"MR":    fmt.Sprintf("%s/src/memos/mr", root),
		"WARNO": fmt.Sprintf("%s/src/orders/warno", root),
		"OPORD": fmt.Sprintf("%s/src/orders/opord", root),
		"FRAGO": fmt.Sprintf("%s/src/orders/frago", root),
	})
	viper.SetDefault("templates", map[string]string{
		"LATEX": fmt.Sprintf("%s/templates/latex", root),
	})
	viper.SetDefault("pdf", map[string]string{
		"MR":    fmt.Sprintf("%s/pdf/memos/mr", root),
		"WARNO": fmt.Sprintf("%s/pdf/orders/warno", root),
		"OPORD": fmt.Sprintf("%s/pdf/orders/opord", root),
		"FRAGO": fmt.Sprintf("%s/pdf/orders/frago", root),
	})
}

func initConfig(home string) {
	path := fmt.Sprintf("%s/.config/eagle-rock/", home)

	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/.config/eagle-rock/", home))
	viper.SetConfigName("documents")
	err = viper.SafeWriteConfig()
	if err != nil {
		if !strings.Contains(err.Error(), "Already Exists") {
			panic(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read configuration: %s\n", err)
	}
}

func initRepository() error {
	repoType := viper.GetString("repo_type")

	switch repoType {
	case "filesystem":
		viper.Set("repo", filesystem.NewFileRepository(
			viper.Get("src").(map[string]string),
			viper.Get("templates").(map[string]string),
			viper.Get("pdf").(map[string]string),
		))
		return nil
	default:
		return errors.Errorf("unsupported repository type: %s", viper.GetString("repo_type"))
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	home, err := homedir.Dir()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to find user's home directory.")
	}

	setDefaultConfig(home)
	initConfig(home)
	err = initRepository()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to initialize repository.")
	}

	cmd.Execute()
}
