package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/golib/color"
)

type Receiver struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Receiver) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("receiver", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Receiver) Synopsis() string {
	return fmt.Sprintf("(%s) Start receiving alarms upon which to start a new workflow execution.", color.Red("starter"))
}

func (this *Receiver) Help() string {
	help := fmt.Sprintf(`
Usage: %s receiver [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
