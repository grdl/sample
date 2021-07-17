package sample

import (
	"fmt"
	"sample/cfg"
)

type Sample struct {
	config *cfg.Config
}

func New(config *cfg.Config) *Sample {
	return &Sample{
		config: config,
	}
}

func (s *Sample) Run() error {
	fmt.Println("hello world")
	fmt.Println(s.config)
	return nil
	//return fmt.Errorf("wot")
}
