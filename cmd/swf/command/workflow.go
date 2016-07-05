package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Workflow struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Workflow) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("workflow", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Workflow) Synopsis() string {
	return "Register/List/Manipulate workflow and workflow type."
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
