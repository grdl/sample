package sample

import (
	"fmt"
)

type Config struct {
	LogLevel string
}

type Sample struct {
	config *Config
}

func New(config *Config) *Sample {
	return &Sample{
		config: config,
	}
}

func (s *Sample) Run() error {
	fmt.Println("hello world")
	fmt.Println(s.config)
	return fmt.Errorf("wot")
	//return nil
}
