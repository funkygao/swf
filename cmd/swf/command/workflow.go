package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gafka/ctx"
	"github.com/funkygao/gocli"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/sdk/go/v1"
)

type Workflow struct {
	Ui  cli.Ui
	Cmd string

	listMode                    bool
	zone, cluster               string
	domain, regName, regVersion string
}

func (this *Workflow) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("workflow", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	cmdFlags.StringVar(&this.regName, "register", "", "")
	cmdFlags.StringVar(&this.domain, "domain", "", "")
	cmdFlags.StringVar(&this.regVersion, "version", "v1", "")
	cmdFlags.BoolVar(&this.listMode, "list", false, "")
	cmdFlags.StringVar(&this.zone, "z", ctx.DefaultZone(), "")
	cmdFlags.StringVar(&this.cluster, "c", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if this.listMode {
		this.listWorkflowTypes()
		return
	}

	if this.regName != "" {
		this.registerWorkflowType()
		return
	}

	return
}

func (this *Workflow) registerWorkflowType() {
	input := &models.RegisterWorkflowTypeInput{}
	input.Cluster = this.cluster
	input.Name = this.regName
	input.Version = this.regVersion
	input.Domain = this.domain

	_, err := swfapi.WithZone(this.zone).RegisterWorkflowType(input)
	if err != nil {
		this.Ui.Error(err.Error())
		return
	}

	this.Ui.Info(fmt.Sprintf("workflow type %+v registered", input))
}

func (this *Workflow) listWorkflowTypes() {
}

func (*Workflow) Synopsis() string {
	return "Register/List/Modify workflow type."
}

func (this *Workflow) Help() string {
	help := fmt.Sprintf(`
Usage: %s workflow -z <zone> -c <cluster> [options]

    %s

    -register <name>
      Register a new workflow type.

      -domain <name>

      -version <version>
        The version of the workflow type.

        NOTE:
            The workflow type consists of name and version, the combination
            of which must be unique within the app.

    -list
      Returns information about workflow types in the specified app.

      [-name <name>]

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
