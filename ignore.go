package main

import (
	"io/ioutil"
	"os"
)
// GitRepo git ginore repo:https://github.com/github/gitignore 
// local path
type GitRepo struct {
	Dir string
	Url string
}

// NewGitRepo create local gitignore repo
func NewGitRepo(dir string, url string)(*GitRepo,error) {
	println(dir)
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
			if file.Name() == ".git" {
				continue
			}
			if file.Name() == ".github" {
				continue
			}
			// 递归
			path := dir+string(os.PathSeparator)+file.Name()
			subPaths , err := readDir(path)
			if err != nil {
				return nil , err
			}
			filePaths = append(filePaths, subPaths...)
			continue
		}
		path := dir+string(os.PathSeparator)+file.Name()
		println(path)
		filePaths = append(filePaths,path)
	}
	return filePaths,err
}
func (git *GitRepo)Sync()error {
	return nil
}