package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Kickoff struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Kickoff) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("kickoff", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Kickoff) Synopsis() string {
	return "Starts an execution of the workflow type using the provided workflowId and input data."
}

func (this *Kickoff) Help() string {
	help := fmt.Sprintf(`
Usage: %s kickoff [options]

    %s

    -workflow-id <value>
      The user defined identifier associated with the workflow execution.
      You cannot have two open workflow executions with the same workflowId at the same time.

    -workflow-type <name,version>
      The type of the workflow to start.

    -in <data>
      The input for the workflow execution.
       This is a free form string which should be meaningful to the workflow you are starting.

    -queue <queue>

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
