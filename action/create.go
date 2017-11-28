package action

import (
	"fmt"
	"github.com/m0cchi/gfalcon/model"
	"github.com/urfave/cli"
)

var CreateFlags = []cli.Flag{}

// Create service of gfalcon
// args[0]: name
func CreateService(c []string) {
	if len(c) < 1 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return

	}
	name := c[0]
	model.CreateService(db, name)
	fmt.Println(model.GetService(db, name))
}

// Create models
func Create(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return
	}
	target := c.Args()[0]
	args := c.Args()[1:]
	switch target {
	case "service":
		CreateService(args)
	default:
		fmt.Println("unknown command")
	}

}
