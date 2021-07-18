package main

import (
	"fmt"
	"os"
	"sample"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:          "sample",
	Short:        "Sample application",
	Version:      sample.Version(),
	RunE:         run,
	SilenceUsage: true, // We don't want to show usage on legit errors
}

// vip is a local instance of Viper available only inside main package.
var vip = viper.New()

func init() {
	cmd.PersistentFlags().BoolP("help", "h", false, "Print this help and exit")
	cmd.PersistentFlags().BoolP("version", "v", false, "Print version and exit")
	cmd.PersistentFlags().StringP("level", "l", "info", "Log level [error, info, debug]")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	vip.AutomaticEnv()
	vip.SetEnvPrefix("SAMPLE")
	vip.AddConfigPath(".")
	vip.SetConfigName(".env")
	vip.SetConfigType("env")

	err := vip.ReadInConfig()
	// Ignore error if config file is not found, default to env vars
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		cobra.CheckErr(err)
	}

	err = vip.BindPFlags(cmd.PersistentFlags())
	cobra.CheckErr(err)
}

func loadAndValidateConfig(args []string) (*sample.Config, error) {
	logLevel := vip.GetString("level")
	validLogLevels := map[string]struct{}{
		"info":  {},
		"error": {},
		"debug": {},
	}

	if _, ok := validLogLevels[logLevel]; !ok {
		return nil, fmt.Errorf("level flag contains invalid value; valid values: %v", validLogLevels)
	}

	config := &sample.Config{
		LogLevel: logLevel,
	}

	return config, nil
}

func run(cmd *cobra.Command, args []string) error {
	config, err := loadAndValidateConfig(args)
	if err != nil {
		return err
	}

	app := sample.New(config)
	return app.Run()
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
