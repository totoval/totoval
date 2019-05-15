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
	return "say:hello-world {hi}"
}

func (hw *HelloWorld) Description() string {
	return "Say Hello"
}

func (hw *HelloWorld) Handler(arg *cmd.Arg) error {
	hi, err := arg.Get("hi")
	if err != nil {
		return err
	}

	if hi == nil {
		fmt.Println("Hello World")
		return nil
	}

	fmt.Println(*hi)
	return nil
}
