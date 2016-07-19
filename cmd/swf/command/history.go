package command

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/swf/models"
	//"github.com/funkygao/swf/sdk/go/v1"
)

type History struct {
	Ui  cli.Ui
	Cmd string

	exec models.WorkflowExecution
}

func (this *History) Run(args []string) (exitCode int) {
	var exec string
	cmdFlags := flag.NewFlagSet("history", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	cmdFlags.StringVar(&exec, "exec", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	parts := strings.SplitN(exec, ",", 2)
	this.exec.WorkflowId = parts[0]
	runId, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {

	}
	if len(parts) > 1 {
		this.exec.RunId = runId
	}

	this.listHistory()

	return
}

func (this *History) listHistory() {
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
