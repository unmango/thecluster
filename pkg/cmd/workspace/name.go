package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	"github.com/unmango/thecluster/pkg/project"
	"github.com/unmango/thecluster/pkg/workspace"
)

func NewName() *cobra.Command {
	return &cobra.Command{
		Use:     "name [PATH]",
		Short:   "Print the name of the workspace",
		Aliases: []string{"n"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			project, err := project.LocalGit(ctx)
			if err != nil {
				cli.Fail(err)
			}

			w, err := workspace.Load(ctx, project, args[0])
			if err != nil {
				cli.Fail(err)
			}

			fmt.Println(w.Name())
		},
	}
}
