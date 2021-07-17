package sample

import "fmt"

type Sample struct {
}

func New() *Sample {
	return &Sample{}
}

func (s *Sample) Run() error {
	fmt.Println("hello world")
	return nil
}
