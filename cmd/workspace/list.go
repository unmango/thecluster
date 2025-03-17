package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	"github.com/unmango/thecluster/pkg/project"
	"github.com/unmango/thecluster/pkg/workspace"
)

func NewList() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List workspaces in the current git context",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			project, err := project.LocalGit(ctx)
			if err != nil {
				cli.Fail(err)
			}

			ws, err := workspace.List(ctx, project)
			if err != nil {
				cli.Fail(err)
			}

			for w := range ws {
				fmt.Println(w)
			}
		},
	}
}
