package main

import (
	"os"

	"github.com/funkygao/gocli"
	"github.com/funkygao/swf/apps/alarmer/command"
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
		"receiver": func() (cli.Command, error) {
			return &command.Receiver{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"sms": func() (cli.Command, error) {
			return &command.Sms{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"marker": func() (cli.Command, error) {
			return &command.Marker{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"escalate": func() (cli.Command, error) {
			return &command.Escalate{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},

		"phone": func() (cli.Command, error) {
			return &command.Phone{
				Ui:  ui,
				Cmd: cmd,
			}, nil
		},
	}

}
