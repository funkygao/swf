package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/swf-go/v1"
)

type Workflow struct {
	Ui  cli.Ui
	Cmd string

	workflowType models.WorkflowType
}

func (this *Workflow) Run(args []string) (exitCode int) {
	var (
		listMode, registerMode bool
	)
	cmdFlags := flag.NewFlagSet("workflow", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	cmdFlags.BoolVar(&listMode, "list", true, "")
	cmdFlags.BoolVar(&registerMode, "register", false, "")
	cmdFlags.StringVar(&this.workflowType.Name, "name", "", "")
	cmdFlags.StringVar(&this.workflowType.Version, "version", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	switch {
	case listMode:
		this.listWorkflowTypes()

	case registerMode:
		this.registerWorkflowType()
	}

	return
}

func (this *Workflow) registerWorkflowType() {
	input := &swfapi.RegisterWorkflowTypeInput{
		WorkflowType: this.workflowType,
	}
	_, err := swfapi.Default().RegisterWorkflowType(input)
	if err != nil {
		this.Ui.Error(err.Error())
		return
	}

	this.Ui.Info("registered")
}

func (this *Workflow) listWorkflowTypes() {
	swfapi.Default().ListWorkflowTypes()
}

func (*Workflow) Synopsis() string {
	return "Register/List/Modify workflow and workflow type."
}

func (this *Workflow) Help() string {
	help := fmt.Sprintf(`
Usage: %s workflow [options]

    %s

    -register <name>
      Register a new workflow type.

      -version <version>
        The version of the workflow type.

        NOTE:
            The workflow type consists of name and version, the combination
            of which must be unique within the app.

      [-description <value>]

    -list
      Returns information about workflow types in the specified app.

      [-name <name>]

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
