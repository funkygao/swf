package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Generate struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Generate) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("generate", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Generate) Synopsis() string {
	return "Generate boilerplate code of decider or worker."
}

func (this *Generate) Help() string {
	help := fmt.Sprintf(`
Usage: %s generate [options]

    %s

    -t <decider|worker>

    -lang <go|php|java>

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
