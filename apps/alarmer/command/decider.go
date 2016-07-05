package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Decider struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Decider) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("decider", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Decider) Synopsis() string {
	return "(decider) Start the decider process."
}

func (this *Decider) Help() string {
	help := fmt.Sprintf(`
Usage: %s decider [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
