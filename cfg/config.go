package cfg

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Version metadata set by ldflags during the build.
var (
	version string
	commit  string
	date    string
)

// vip is a local instance of Viper available only inside cfg package.
var vip *viper.Viper

// Version returns a string with version metadata: version number, git sha and build date.
// It returns "development" if version variables are not set during the build.
func Version() string {
	if version == "" {
		return "development"
	}

	return fmt.Sprintf("%s - revision %s built at %s", version, commit[:6], date)
}

type Config struct {
	LogLevel string
}

// Load returns a Config populated with values from flags, env variables and config file.
// If config can't be loaded or values are invalid, an error is returned.
func Load(cmd *cobra.Command, args []string) (*Config, error) {
	err := initViper(cmd)
	if err != nil {
		return nil, err
	}

	return loadAndValidateConfig(args)
}

func initViper(cmd *cobra.Command) error {
	vip = viper.New()

	vip.AutomaticEnv()
	vip.SetEnvPrefix("SAMPLE")
	vip.AddConfigPath(".")
	vip.SetConfigName(".env")
	vip.SetConfigType("env")

	err := vip.ReadInConfig()
	// Ignore error if config file is not found, default to env vars
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return err
	}

	err = vip.BindPFlags(cmd.PersistentFlags())
	return err
}

func loadAndValidateConfig(args []string) (*Config, error) {
	logLevel := vip.GetString("level")
	validLogLevels := map[string]struct{}{
		"info":  {},
		"error": {},
		"debug": {},
	}

	if _, ok := validLogLevels[logLevel]; !ok {
		return nil, fmt.Errorf("level flag contains invalid value; valid values: %v", validLogLevels)
	}

	config := &Config{
		LogLevel: logLevel,
	}

	return config, nil
}
