package main

import (
	"os"

	"github.com/funkygao/gocli"
	"github.com/funkygao/swf/cmd/swf/command"
)

var commands map[string]cli.CommandFactory

func init() {
	ui := &cli.ColoredUi{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			Reader:      os.Stdin,
			ErrorWriter: os.Stderr,
		},
		InfoColor:  cli.UiColorGreen,
		ErrorColor: cli.UiColorRed,
		WarnColor:  cli.UiColorYellow,
	}

	cmd := os.Args[0]
	commands = map[string]cli.CommandFactory{
		"workflow": func() (cli.Command, error) {
			return &command.Workflow{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"?": func() (cli.Command, error) {
			return &command.FAQ{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"activity": func() (cli.Command, error) {
			return &command.Activity{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"generate": func() (cli.Command, error) {
			return &command.Generate{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"kickoff": func() (cli.Command, error) {
			return &command.Kickoff{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"history": func() (cli.Command, error) {
			return &command.History{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},
	}

}
