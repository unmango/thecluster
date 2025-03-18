package main

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	"github.com/unmango/thecluster/app"
	"github.com/unmango/thecluster/project"
)

func runApp(ctx context.Context) {
	p := tea.NewProgram(&app.Model{},
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

var root = &cobra.Command{
	Use:   "thecluster",
	Short: "Manage your CLUSTER with style",
	Run: func(cmd *cobra.Command, args []string) {
		// https://github.com/charmbracelet/bubbletea/issues/860
		if isatty.IsTerminal(os.Stdout.Fd()) {
			runApp(cmd.Context())
		} else {
			run(cmd.Context())
		}
	},
}

func main() {
	if err := root.Execute(); err != nil {
		cli.Fail(err)
	}
}
