package pipefire_interface

import (
    "fmt"
)
//PipefirePlugin Available operations for the plugin
type PipefirePlugin interface {
	GetVersion() string
	GetName() string
	IsPipeline() bool
	IsExtention() bool
}

//Pipeline is the
type Pipeline interface {
	StartListener(listenerError chan error)
	Execute(string) []error
	Close() error
}

//FooBar
func FooBar() {
  fmt.Print("FooBar")
}
