package main

import (
	"github.com/spf13/cobra"
	"os"
	"sample"
)

var sampleCmd = &cobra.Command{
	Use:          "sample",
	Short:        "Sample application",
	Version:      "0.0.0",
	RunE:         run,
	SilenceUsage: true, // We don't want to show usage on legit errors
}

func run(cmd *cobra.Command, args []string) error {
	app := sample.New()
	return app.Run()
}

func main() {
	if err := sampleCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
