package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unmango/thecluster/pkg/cmd/workspace"
)

func NewWorkspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workspace",
		Short:   "Manage local git workspaces",
		Aliases: []string{"ws"},
	}
	cmd.AddCommand(
		workspace.NewList(),
	)

	return cmd
}
