package command

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type FAQ struct {
	Ui  cli.Ui
	Cmd string
}

func (this *FAQ) Run(args []string) (exitCode int) {
	content := fmt.Sprintf(`
How to use SWF?
    swf workflow -register <name> -version <version> -z <zone> -c <cluster>
    swf workflow -list -z <zone> -c <cluster>
    swf activity -register <name> -version <version> -z <zone> -c <cluster>
    swf kickoff -z <zone> -c <cluster> -workflow-type <name,version> -workflow-id <id> -input <data>
    swf history -z <zone> -c <cluster>

How can I get help?
    %s <command> -h

`, this.Cmd)

	this.Ui.Output(strings.TrimSpace(content))

	return
}

func (*FAQ) Synopsis() string {
	return "FAQ."
}

func (this *FAQ) Help() string {
	help := fmt.Sprintf(`
Usage: %s ?

    FAQ

`, this.Cmd)
	return strings.TrimSpace(help)
}
