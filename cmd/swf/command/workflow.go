package command

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Workflow struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Workflow) Run(args []string) (exitCode int) {
	return
}

func (*Workflow) Synopsis() string {
	return "Register/List/Manipulate workflow and workflow type"
}

func (this *Workflow) Help() string {
	help := fmt.Sprintf(`
Usage: %s workflow [options]

    Register/List/Manipulate workflow and workflow type

`, this.Cmd)
	return strings.TrimSpace(help)
}
