package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg/context"
	"github.com/unmango/thecluster/pkg/version"
)

func NewVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version [NAME]",
		Short:   "Print the version of the specified dependency",
		Long:    "Print the version of the specified dependency, or the version of this tool if no name is provided",
		Aliases: []string{"v", "ver"},
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, err := context.LocalRepo(cmd.Context())
			if err != nil {
				internal.Fail(err)
			}

			if len(args) == 1 {
				version, err := version.Get(ctx, args[0])
				if err != nil {
					internal.Fail(err)
				}

				fmt.Println(version)
			} else {
				fmt.Println("TODO")
			}
		},
	}

	return cmd
}
