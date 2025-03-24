package main

import (
	"github.com/unmango/go/cli"
	"github.com/unmango/thecluster/cmd"
)

var root = cmd.New()

func main() {
	if err := root.Execute(); err != nil {
		cli.Fail(err)
	}
}
