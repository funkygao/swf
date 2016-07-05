package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Escalate struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Escalate) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("escalate", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Escalate) Synopsis() string {
	return "(worker) Escalate this alarm to a higher level."
}

func (this *Escalate) Help() string {
	help := fmt.Sprintf(`
Usage: %s escalate [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
