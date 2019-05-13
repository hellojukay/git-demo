package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	prompt "github.com/c-bata/go-prompt"
	"github.com/fatih/color"
)

var IGNORE_HOME string
var app *cli.App

func init() {
	app = cli.NewApp()
	app.Usage = "manager git ignore file"
	app.Commands = []cli.Command{
		{
			Name:         "init",
			Aliases:      []string{"i"},
			Usage:        "init git ignore file",
			Action:       initIgnore,
		},
		{
			Name:   "reset",
			Usage:  "clean git ignore file",
			Action: func(ctx *cli.Context) {
				cleanFile(".gitignore")
				color.Green("%s \n", "SUCCESS")
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		cli.ShowCommandHelpAndExit(c, "", 1)
		fmt.Fprintf(c.App.Writer, app.ArgsUsage)
		return nil
	}
}

func initIgnore(c *cli.Context) {
	var home, _ = os.UserHomeDir()
	git, err := NewGitRepo(fmt.Sprintf("%s/.gitignore", home), "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		cli.OsExiter(1)
	}
	files, err := git.Files()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		cli.OsExiter(1)
	}
	t := func()string {
		var s []prompt.Suggest
		for _, file := range files {
			s = append(s,prompt.Suggest{
				Text:filepath.Base(file),
				Description:file,
			})
		}
		return prompt.Input(">", func(document prompt.Document) []prompt.Suggest {
			return prompt.FilterHasPrefix(s, document.GetWordBeforeCursor(), true)
		})
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		cli.OsExiter(1)
	}
	content , err := git.ReadFile(t)
	if err != nil {
		fmt.Fprintf(os.Stderr,"%s\n",err)
		cli.OsExiter(1)
	}
	appendIgnore(content)
	println(content)
	color.Green("%s\n","SUCCESS")
}


func main() {
	app.Run(os.Args)
}
