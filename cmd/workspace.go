package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unmango/thecluster/cmd/workspace"
)

func NewWorkspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workspace",
		Short:   "Manage local git workspaces",
		Aliases: []string{"ws"},
	}
	cmd.AddCommand(
		workspace.NewInstall(),
		workspace.NewList(),
		workspace.NewName(),
	)

	return cmd
}
