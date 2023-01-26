package main

import "testing"

// t is the test handler
func TestNewDeck(t *testing.T) {
	myDeck := newDeck()

	if len(myDeck) != 16 {
		t.Errorf("Expect 16 cards, got %v", len(myDeck))
	}
}

/*
test function name follows pascal case
*/

/* run test from terminal
➜  learning-golang git:(master) ✗ go test
go: cannot find main module, but found .git/config in /Users/hakw/code/learning-golang
        to create a module there, run:
        go mod init

GO111MODULE=on

➜  learning-golang git:(master) ✗ go env -w GO111MODULE=auto
warning: go env -w GO111MODULE=... does not override conflicting OS environment variable

set export GO111MODULE=auto in ~/.zsh doesn't work

set export GO111MODULE=auto in terminal session works temporarily

go init mod learning-golang solved the problem
*/

/* run test from pycharm

GOROOT=/opt/homebrew/opt/go/libexec #gosetup
GOPATH=/Users/hakw/go #gosetup
/opt/homebrew/opt/go/libexec/bin/go test -c -o /private/var/folders/kz/w9khlcc541s7gl8w2m387dk40000gp/T/GoLand/___TestNewDeck_in_deck_test_go.test /Users/hakw/code/learning-golang/deck_test.go #gosetup
# command-line-arguments [command-line-arguments.test]
./deck_test.go:7:12: undefined: newDeck

Compilation finished with exit code 2

go init mod learning-golang solved the problem
*/
