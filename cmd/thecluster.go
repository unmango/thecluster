package cmd

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
	"github.com/unmango/thecluster/app"
	"github.com/unmango/thecluster/project"
)

type Options struct {
	work.ChdirOptions
}

func New() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "thecluster",
		Short: "Manage your CLUSTER with style",
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.Chdir(cmd.Context()); err != nil {
				cli.Fail(err)
			}

			// https://github.com/charmbracelet/bubbletea/issues/860
			if isatty.IsTerminal(os.Stdout.Fd()) {
				runApp(cmd.Context())
			} else {
				run(cmd.Context())
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")

	return cmd
}

func runApp(ctx context.Context) {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			cli.Fail(err)
		}
		defer f.Close()
	}

	p := tea.NewProgram(app.New(),
		tea.WithContext(ctx),
	)
	if _, err := p.Run(); err != nil {
		cli.Fail(err)
	}
}

func run(ctx context.Context) {
	if proj, err := project.Load(ctx); err != nil {
		cli.Fail(err)
	} else {
		fmt.Println("Project:", proj.Dir.Path())
	}
}
