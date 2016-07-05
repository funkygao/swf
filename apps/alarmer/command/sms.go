package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/golib/color"
)

type Sms struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Sms) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("sms", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Sms) Synopsis() string {
	return fmt.Sprintf("(%s) Send sms to person in charge of the alarm.", color.Blue("worker"))
}

func (this *Sms) Help() string {
	help := fmt.Sprintf(`
Usage: %s sms [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
