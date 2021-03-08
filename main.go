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
	"path/filepath"
	"strings"
)

const BaseVersion = "0.1.0-develop"
var BuildTime string
var GitRevision string
var GitBranch string

func setDefaultConfig(home string) {
	root := filepath.Join(home, "Documents/Org/Policy")
	viper.SetDefault("repo_type", "filesystem")
	viper.SetDefault("src", map[string]string{
		"MR":    filepath.Join(root, "src/memos/mrs"),
		"WARNO": filepath.Join(root, "src/orders/warnos"),
		"OPORD": filepath.Join(root, "src/orders/opords"),
		"FRAGO": filepath.Join(root, "src/orders/fragos"),
	})
	viper.SetDefault("templates", map[string]string{
		"LATEX": filepath.Join(root, "src/templates/latex"),
	})
	viper.SetDefault("pdf", map[string]string{
		"MR":    filepath.Join(root, "pdf/memos/mrs"),
		"WARNO": filepath.Join(root, "pdf/orders/warnos"),
		"OPORD": filepath.Join(root, "pdf/orders/opords"),
		"FRAGO": filepath.Join(root, "pdf/orders/fragos"),
	})
	viper.SetDefault(
		"texinputs",
		fmt.Sprintf(
			"%s:%s",
			filepath.Join(root, "src/assets"),
			filepath.Join(root, "src/assets/latex"),
		),
	)
}

func initConfig(home string) {
	path := fmt.Sprintf("%s/.config/eagle-rock/", home)

	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/.config/eagle-rock/", home))
	viper.SetConfigName("lib")
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
			viper.GetStringMapString("src"),
			viper.GetStringMapString("templates"),
			viper.GetStringMapString("pdf"),
		))
		return nil
	default:
		return errors.Errorf("unsupported repository type: %s", viper.GetString("repo_type"))
	}
}

func main() {
	version := fmt.Sprintf(
		"%s+%s.%s.%s",
		BaseVersion,
		GitBranch,
		GitRevision,
		BuildTime,
	)
	viper.Set("VERSION", version)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	//zerolog.SetGlobalLevel(zerolog.DebugLevel)
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
