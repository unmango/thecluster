package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg/context"
	"github.com/unmango/thecluster/pkg/workspace"
)

func NewName() *cobra.Command {
	return &cobra.Command{
		Use:     "name [PATH]",
		Short:   "Print the name of the workspace",
		Aliases: []string{"n"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, err := context.LocalRepo(cmd.Context())
			if err != nil {
				internal.Fail(err)
			}

			w, err := workspace.Load(ctx, args[0])
			if err != nil {
				internal.Fail(err)
			}

			fmt.Println(w.Name())
		},
	}
}
