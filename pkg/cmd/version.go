package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	"github.com/unmango/thecluster/pkg/project"
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
			ctx := cmd.Context()
			project, err := project.LocalGit(ctx)
			if err != nil {
				cli.Fail(err)
			}

			if len(args) == 1 {
				version, err := version.Get(project, args[0])
				if err != nil {
					cli.Fail(err)
				}

				fmt.Println(version)
			} else {
				fmt.Println("TODO")
			}
		},
	}

	return cmd
}
