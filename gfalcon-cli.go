package main

import (
	"fmt"
	"github.com/m0cchi/gfalcon-cli/action"
	"github.com/urfave/cli"
	"os"
)

func main() {
	err := action.Init()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	defer action.Final()
	app := cli.NewApp()
	app.Name = "gfalcon-cli"
	app.Usage = "control models of gfalcon"
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create models",
			Action:  action.Create,
			//Flags:   action.CreateFlags,
		},
	}
	app.Run(os.Args)
}
