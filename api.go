package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)
func getItems()[]string{
	response ,err := http.Get("https://gitignore.io/api/list")
	if err != nil {
		fmt.Errorf("get ignore list error , network error %s",err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	body , err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("get ignore list error , network error %s",err.Error())
		os.Exit(1)
	}
	return strings.Split(string(body),",")
}
