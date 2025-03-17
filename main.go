package main

import (
	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
)

var root = &cobra.Command{
	Use:   "thecluster",
	Short: "Manage your CLUSTER",
}

func main() {
	if err := root.Execute(); err != nil {
		cli.Fail(err)
	}
}
