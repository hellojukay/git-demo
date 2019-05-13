package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GitRepo git ginore repo:https://github.com/github/gitignore
// local path
type GitRepo struct {
	Dir string
	Url string
}

// NewGitRepo create local gitignore repo
func NewGitRepo(dir string, url string) (*GitRepo, error) {
	// check update
	return &GitRepo{
		Dir: dir,
		Url: url,
	}, nil
}
func (git *GitRepo) Files() ([]string, error) {
	return readDir(git.Dir)
}
func readDir(dir string) ([]string, error) {
	var filePaths []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			// 递归
			path := dir + string(os.PathSeparator) + file.Name()
			subPaths, err := readDir(path)
			if err != nil {
				return nil, err
			}
			filePaths = append(filePaths, subPaths...)
			continue
		}
		path := dir + string(os.PathSeparator) + file.Name()
		filePaths = append(filePaths, path)
	}
	return filter(filePaths), err
}
func filter(paths []string) []string {
	var s []string
	var suffix = "gitignore"
	for _, path := range paths {
		if strings.HasSuffix(path, suffix) {
			s = append(s, path)
		}
	}
	return s
}
func (git *GitRepo) Sync() error {
	return nil
}

func cleanFile(path string) {
	fh , err := os.OpenFile(path,os.O_TRUNC,0666)
	if err != nil {
		fmt.Fprintf(os.Stderr,"%s\n",err)
		os.Exit(1)
	}
	fh.Close()
}
