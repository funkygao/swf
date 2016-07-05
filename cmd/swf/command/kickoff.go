package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/swf-go/v1"
)

type Kickoff struct {
	Ui  cli.Ui
	Cmd string

	workflowId   string
	input        string
	workflowType models.WorkflowType
}

func (this *Kickoff) Run(args []string) (exitCode int) {
	var workflowType string
	cmdFlags := flag.NewFlagSet("kickoff", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	cmdFlags.StringVar(&this.workflowId, "workflow-id", "", "")
	cmdFlags.StringVar(&this.input, "input", "", "")
	cmdFlags.StringVar(&workflowType, "workflow-type", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	parts := strings.SplitN(workflowType, ",", 2)
	this.workflowType.Name = parts[0]
	if len(parts) > 1 {
		this.workflowType.Version = parts[1]
	}

	this.startExecution()

	return
}

func (this *Kickoff) startExecution() {
	swfapi.Default().StartWorkflowExecution()
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

    -input <data>
      The input for the workflow execution.
      This is a free form string which should be meaningful to the workflow you are starting.

    -queue <queue>

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
