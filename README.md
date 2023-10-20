# README

Study notes for Udemy course Go: The Complete Developer's Guide (Golang).

## Commands
`go mod init [module-path]`
declare a Go module
module-path is the global unique identifier, case sensitive

`go mod edit -module [module-path]`
rename a module

`gonew github.com/Magicbeanbuyer/learning-golang/module_template github.com/Magicbeanbuyer/learning-golang/<NAME>`
create a new module from template.

## Concepts

Module
declared with `go.mod`
has a global unique identifier
everything inside a module are version together

Package

a folder (and subfolders) of `.go` files

one package per directory

package and dir don't have to have the same name, import dir name, and call by package name https://golangbyexample.com/package-folder-name-golang/

A package's exported constants, variables and functions all start with an uppercase letter in their name

## Envs

`GOPATH`

third party tools path

`go install` installs modules to `GOPATH`
exported via .zshrc

`GOPATH`
no need to export explicitly anymore
path of Go development environment