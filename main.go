package main

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/cmd"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func setDefaultConfig(home string) {
	viper.SetDefault("directory", fmt.Sprintf("%s/Documents/Org/Policy", home))
	viper.SetDefault("index", fmt.Sprintf("%s/Documents/Org/Policy/content/doc-index.yml", home))
	viper.SetDefault("content", map[string]string{
		"mr_dir":    fmt.Sprintf("%s/Documents/Org/Policy/content/mr", home),
		"warno_dir": fmt.Sprintf("%s/Documents/Org/Policy/content/warno", home),
		"opord_dir": fmt.Sprintf("%s/Documents/Org/Policy/content/opord", home),
		"frago_dir": fmt.Sprintf("%s/Documents/Org/Policy/content/frago", home),
	})
	viper.SetDefault("published", map[string]string{
		"mr_dir":    fmt.Sprintf("%s/Documents/Org/Policy/mr", home),
		"warno_dir": fmt.Sprintf("%s/Documents/Org/Policy/warno", home),
		"opord_dir": fmt.Sprintf("%s/Documents/Org/Policy/opord", home),
		"frago_dir": fmt.Sprintf("%s/Documents/Org/Policy/frago", home),
	})
	viper.SetDefault("latex_templates", map[string]string{
		"mr":    fmt.Sprintf("%s/Documents/Org/Policy/templates/latex/mr.template", home),
		"warno": fmt.Sprintf("%s/Documents/Org/Policy/templates/latex/warno.template", home),
		"opord": fmt.Sprintf("%s/Documents/Org/Policy/templates/latex/opord.template", home),
		"frago": fmt.Sprintf("%s/Documents/Org/Policy/templates/latex/frago.template", home),
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

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	home, err := homedir.Dir()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to find user's home directory.")
	}

	setDefaultConfig(home)
	initConfig(home)

	cmd.Execute()
}
