package action

import (
	"fmt"
	"github.com/m0cchi/gfalcon/complex"
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

// args[0]: teamID
// args[1]: userID
// args[2]: password
func CreateUser(c []string) {
	if len(c) < 3 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return
	}
	teamID := c[0]
	userID := c[1]
	password := c[2]
	team, err := model.GetTeam(db, teamID)
	if err != nil {
		fmt.Printf("failed to get team: %v\n", err)
		return
	}

	fmt.Println(complex.CreateUser(db, team.IID, userID, password))
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
	case "user":
		CreateUser(args)
	default:
		fmt.Println("unknown command")
	}
}

// Create models
func Create(c *cli.Context) {
	create(c.Args())
}
