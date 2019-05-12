package main

import (
	"io/ioutil"
	 "path/filepath"

)
// GitRepo git ginore repo:https://github.com/github/gitignore 
// local path
type GitRepo struct {
	Dir string
	Url string
}

// NewGitRepo create local gitignore repo
func NewGitRepo(dir string, url string)(*GitRepo,error) {
	// check update
	return & GitRepo{
		Dir:dir,
		Url:url,
	},nil
}
func (git *GitRepo)Files()([]string,error) {
	return readDir(git.Dir)
}
func readDir(dir string)([]string,error) {
	var filePaths []string
	files , err := ioutil.ReadDir(dir)
	if err != nil {
		return nil ,err
	}
	for _, file := range files {
		if file.IsDir() {
			// 递归
			subDir , _ := filepath.Abs(file.Name())
			subPaths , err := readDir(subDir)
			if err != nil {
				return nil , err
			}
			filePaths = append(filePaths, subPaths...)
		}
		path ,_ := filepath.Abs(file.Name())
		filePaths = append(filePaths,path)
	}
	return filePaths,err
}
func (git *GitRepo)Sync()error {
	return nil
}