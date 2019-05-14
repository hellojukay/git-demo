package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// GitRepo git ginore template repo
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

func (git *GitRepo) ReadFile(fileName string) (string, error) {
	path := git.Dir + string(os.PathSeparator) + "templates" + string(os.PathSeparator) + fileName
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", nil
	}
	return string(bytes), err
}
func appendIgnore(content string) error {
	fh, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = fh.WriteString(fmt.Sprintf("\n#%s \n", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		return err
	}
	_, err = fh.WriteString(fmt.Sprintf("%s\n", content))
	if err != nil {
		return err
	}
	return nil
}
func cleanFile(path string) {
	fh, err := os.OpenFile(path, os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fh.Close()
}
