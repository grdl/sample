package main

import (
	"sample"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:          "sample",
	Short:        "Sample application",
	Version:      sample.Version(),
	Run:          run,
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

func run(cmd *cobra.Command, args []string) {
	config := &sample.Config{
		LogLevel: vip.GetString("level"),
	}

	app, err := sample.New(config)
	cobra.CheckErr(err)

	cobra.CheckErr(app.Run())
}

func main() {
	cobra.CheckErr(cmd.Execute())
}
