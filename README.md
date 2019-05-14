# git-ignore
This is a git plugin to add file or dir to gitignore file

[![Build Status](https://travis-ci.org/hellojukay/git-ignore.svg?branch=master)](https://travis-ci.org/hellojukay/git-ignore)

## How to use it
> make sure $GOPATH/bin in PATH

install with golang
```shell
go install github.com/hellojukay/git-ignore
git clone git@github.com:dvcs/gitignore.git tmp && mv tmp $HOME/.gitignore && rm -rf tmp
```
add a git ignore template file
```shell
// git ignore init
git ignore i
```
![init](init.gif)
clean git ignore file
```shell
git ignore clean
```
![clean](clean.gif)
