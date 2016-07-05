package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/golib/color"
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
	return fmt.Sprintf("(%s) Start the decider process.", color.Yellow("decider"))
}

func (this *Decider) Help() string {
	help := fmt.Sprintf(`
Usage: %s decider [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
