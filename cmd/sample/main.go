package main

import (
	"github.com/spf13/cobra"
	"os"
	"sample"
	"sample/cfg"
)

var cmd = &cobra.Command{
	Use:          "sample",
	Short:        "Sample application",
	Version:      cfg.Version(),
	RunE:         run,
	SilenceUsage: true, // We don't want to show usage on legit errors
}

func init() {
	cmd.PersistentFlags().BoolP("help", "h", false, "Print this help and exit")
	cmd.PersistentFlags().BoolP("version", "v", false, "Print version and exit")
	cmd.PersistentFlags().StringP("level", "l", "info", "Log level [error, info, debug]")
}

func run(cmd *cobra.Command, args []string) error {
	config, err := cfg.Load(cmd, args)
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
