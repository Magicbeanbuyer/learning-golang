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

`go list -m -versions github.com/shopspring/decimal`

list versions of a module

`go install <REPO_URL>@<VERSION>`

download packages to `GOPATH/PKG/mod`, a.k.a `GOMODCACHE`
compile source code into an executable and install it to `GOPATH/bin`

`go get <REPO_URL>@<VERSION>`

install packages to `~/go/pkg/mod` a.k.a `GOMODCACHE`
add modules to `go.mod` & `go.sum`

`go mod tidy`
remove unimported modules from `go.mod` & `go.sum`

`go clean -modcache`

remove dir `~/go/pkg/mod` a.k.a `GOMODCACHE`

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
`go install` installs modules to `GOPATH/bin`
exported via .zshrc

`GOROOT`

path of Go development environment
no need to export explicitly anymore