package main

import (
	"flag"
	"gopkg.in/AlecAivazis/survey.v1"
	"strings"
)

// is init ?
var i bool

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	flag.BoolVar(&i, "i", false, "init and reset gitignore file")
}
func main() {

	// TODO 获取 gitignore list
	// TODO gitignore list to array
	// parse args
	// switch
	// init
	var items = getItems()
	var values []string
	promt := survey.MultiSelect{
		Message: "Choose items:",
		Options: items,
		FilterFn:filter,
		PageSize:5,
		}
		survey.AskOne(&promt,&values,nil)
}

func filter(filter string, options []string)[]string  {
	if len(options) == 0 {
		return options
	}
	var values []string
	for index := range options {
		if strings.Contains(options[index],filter) {
			values = append(values,options[index])
		}
	}
	return values
}