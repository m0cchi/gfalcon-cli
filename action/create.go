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

	fmt.Println(model.CreateService(db, name))

}

// args[0]: name
func CreateTeam(c []string) {
	if len(c) < 1 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return
	}
	name := c[0]

	fmt.Println(model.CreateTeam(db, name))
}

func create(args []string) {
	if len(args) < 1 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return
	}
	target := args[0]
	args = args[1:]
	switch target {
	case "service":
		CreateService(args)
	case "team":
		CreateTeam(args)
	default:
		fmt.Println("unknown command")
	}
}

// Create models
func Create(c *cli.Context) {
	create(c.Args())
}
