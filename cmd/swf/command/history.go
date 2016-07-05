package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type History struct {
	Ui  cli.Ui
	Cmd string
}

func (this *History) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("history", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*History) Synopsis() string {
	return "Returns the history of the specified workflow execution."
}

func (this *History) Help() string {
	help := fmt.Sprintf(`
Usage: %s history [options]

    %s

    -exec <workflowId,runId>
      Specifies the workflow execution for which to return the history.

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
