package action

import (
	"fmt"
	"github.com/m0cchi/gfalcon/complex"
	"github.com/m0cchi/gfalcon/model"
	"github.com/urfave/cli"
)

var CreateFlags = []cli.Flag{}

// Create service of gfalcon
// args[0]: serviceID
func CreateService(c []string) {
	if len(c) < 1 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return

	}
	serviceID := c[0]

	fmt.Println(model.CreateService(db, serviceID))

}

// args[0]: teamID
func CreateTeam(c []string) {
	if len(c) < 1 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return
	}
	teamID := c[0]

	fmt.Println(model.CreateTeam(db, teamID))
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

// args[0]: serviceID
// args[1]: actionID
func CreateAction(c []string) {
	if len(c) < 2 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return
	}

	serviceID := c[0]
	actionID := c[1]
	service, err := model.GetService(db, serviceID)
	if err != nil {
		fmt.Printf("failed to get service: %v\n", err)
		return
	}

	fmt.Println(model.CreateAction(db, service.IID, actionID))
}

// args[0]: serviceID
// args[1]: actionID
// args[2]: teamID
// args[3]: userID
func CreateActionLink(c []string) {
	if len(c) < 4 {
		fmt.Println("gfalcon-cli create [target] [args...]")
		return
	}

	serviceID := c[0]
	actionID := c[1]
	teamID := c[2]
	userID := c[3]

	service, err := model.GetService(db, serviceID)
	if err != nil {
		fmt.Printf("failed to get service: %v\n", err)
		return
	}

	action, err := model.GetAction(db, service.IID, actionID)
	if err != nil {
		fmt.Printf("failed to get action: %v\n", err)
		return
	}

	team, err := model.GetTeam(db, teamID)
	if err != nil {
		fmt.Printf("failed to get team: %v\n", err)
		return
	}

	user, err := model.GetUser(db, team.IID, userID)
	if err != nil {
		fmt.Printf("failed to get user: %v\n", err)
		return
	}
	fmt.Println(model.CreateActionLink(db, action, user))

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
	case "action":
		CreateAction(args)
	case "action-link":
		CreateActionLink(args)
	default:
		fmt.Println("unknown command")
	}
}

// Create models
func Create(c *cli.Context) {
	create(c.Args())
}
