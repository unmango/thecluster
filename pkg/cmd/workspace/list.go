package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg/context"
	"github.com/unmango/thecluster/pkg/workspace"
)

func NewList() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List workspaces in the current git context",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			ctx, err := context.LocalRepo(cmd.Context())
			if err != nil {
				internal.Fail(err)
			}

			ws, err := workspace.List(ctx)
			if err != nil {
				internal.Fail(err)
			}

			for w := range ws {
				fmt.Println(w)
			}
		},
	}
}
