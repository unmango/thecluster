package workspace

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg/progress"
	"github.com/unmango/thecluster/pkg/project"
	"github.com/unmango/thecluster/pkg/workspace"
)

func NewInstall() *cobra.Command {
	return &cobra.Command{
		Use:     "install [PATH]",
		Short:   "Install workspace dependencies",
		Aliases: []string{"i"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			project, err := project.LocalGit(ctx)
			if err != nil {
				internal.Fail(err)
			}

			w, err := workspace.Load(ctx, project, args[0])
			if err != nil {
				internal.Fail(err)
			}

			sub := progress.WriteTo(workspace.Observe(w), os.Stdout)
			defer sub()

			if err = workspace.Install(ctx, w); err != nil {
				internal.Fail(err)
			}
		},
	}
}
