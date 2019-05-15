package commands

import (
	"fmt"

	"github.com/totoval/framework/cmd"
)

func init() {
	cmd.Add(&HelloWorld{})
}

type HelloWorld struct {
}

func (hw *HelloWorld) Command() string {
	return "say:hello-world {test}"
}

func (hw *HelloWorld) Description() string {
	return "Say Hello"
}

func (hw *HelloWorld) Handler(arg *cmd.Arg) error {
	fmt.Println("Hello World")
	return nil
}
