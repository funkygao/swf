package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/golib/color"
)

type Phone struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Phone) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("phone", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Phone) Synopsis() string {
	return fmt.Sprintf("(%s) Manually phone call the person in charge.", color.Blue("worker"))
}

func (this *Phone) Help() string {
	help := fmt.Sprintf(`
Usage: %s phone [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
