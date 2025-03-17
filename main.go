package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	"github.com/unmango/thecluster/project"
)

var root = &cobra.Command{
	Use:   "thecluster",
	Short: "Manage your CLUSTER with style",
	Run: func(cmd *cobra.Command, args []string) {
		if proj, err := project.Load(cmd.Context()); err != nil {
			cli.Fail(err)
		} else {
			fmt.Println("Project:", proj.Dir.Path())
		}
	},
}

func main() {
	if err := root.Execute(); err != nil {
		cli.Fail(err)
	}
}
