package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
	"github.com/manifoldco/promptui"
)
var IGNORE_HOME string
var app *cli.App
func init() {
	app = cli.NewApp()
	app.Usage = "manager git ignore file"
	app.Commands = []cli.Command{
		{
		  Name:    "init",
		  Aliases: []string{"i"},
		  Usage:   "init git ignore file",
		  Action:  initAutoComplete,
		  BashComplete: initAutoComplete,
		},
		{
		  Name:    "reset",
		  Usage:   "clean git ignore file",
		  Action:  initAutoComplete,
		},
	}
	app.Action = func(c *cli.Context)error {
		cli.ShowCommandHelpAndExit(c,"",1)
		fmt.Fprintf(c.App.Writer,app.ArgsUsage)
		return nil
	}
	app.EnableBashCompletion = true
	app.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer,c.Command.FullName())
	}
}

func initAutoComplete(c *cli.Context)  {
	fmt.Fprintf(c.App.Writer,"Hello \n World")
	var home , _ =os.UserHomeDir()
	git, err := NewGitRepo(fmt.Sprintf("%s/.gitignore",home),"")
	if err != nil {
		fmt.Fprintf(os.Stderr,"%s\n",err.Error())
		cli.OsExiter(1)
	}
	files , err := git.Files()
	if err != nil {
		fmt.Fprintf(os.Stderr,"%s\n",err)
		cli.OsExiter(1)
	}
	prompt := promptui.Select{
		Label: "select ignore template",
		Items: files,
	}

	_, result, err := prompt.Run()
	println(result)
}
func main() {
	app.Run(os.Args)
}