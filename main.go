package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

// is init ?
var i bool
var app *cli.App
func init() {
	app = cli.NewApp()
	app.Commands = []cli.Command{
		{
		  Name:    "init",
		  Aliases: []string{"i"},
		  Usage:   "init git ignore file",
		  Action:  func(c *cli.Context) error {
			fmt.Println("added task: ", c.Args().First())
			return nil
		  },
		  BashComplete: initAutoComplete,
		},
		{
		  Name:    "reset",
		  Usage:   "complete a task on the list",
		  Action:  func(c *cli.Context) error {
			fmt.Println("completed task: ", c.Args().First())
			return nil
		  },
		},
	}
	app.Action = func(c *cli.Context)error {
		println("Hello World")
		return nil
	}
	app.EnableBashCompletion = true
	app.BashComplete = func(c *cli.Context) {
		println("Hello World\n")
	}
}

func initAutoComplete(c *cli.Context)  {
	fmt.Fprintf(c.App.Writer,"Hello \n World")
}
func main() {
	app.Run(os.Args)
}