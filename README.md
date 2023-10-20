# README

Study notes for Udemy course Go: The Complete Developer's Guide (Golang).

## Commands
`go mod init [module-path]`
declare a Go module
module-path is the global unique identifier, case sensitive

`go mod edit -module [module-path]`
rename a module

## Concepts

Module
declared with `go.mod`
has a global unique identifier
everything inside a module are version together

Package
one package per directory
package and dir don't have to have the same name, import dir name, and call by package name https://golangbyexample.com/package-folder-name-golang/
a folder (and subfolders) of `.go` files
A package's exported constants, variables and functions all start with an uppercase letter in their name
