package main

import (
	"github.com/spf13/cobra"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg/cmd"
)

var root = &cobra.Command{
	Use:   "thecluster",
	Short: "Manage your CLUSTER",
}

func main() {
	root.AddCommand(
		cmd.NewVersion(),
		cmd.NewWorkspace(),
	)

	if err := root.Execute(); err != nil {
		internal.Fail(err)
	}
}
