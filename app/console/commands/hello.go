package commands

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/totoval/framework/cmd"
)

func init() {
	cmd.Add(&HelloWorld{})
}

type HelloWorld struct {
}

func (hw *HelloWorld) Command() string {
	return "say:hello-world"
}

func (hw *HelloWorld) Aliases() []string {
	return []string{
		"c",
	}
}

func (hw *HelloWorld) Description() string {
	return "Say Hello"
}

func (hw *HelloWorld) Handler(c *cli.Context) error {
	fmt.Println("Hello World")
	return nil
}
