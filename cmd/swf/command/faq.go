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
    swf workflow -register <name> -version <version>
    swf activity -register <name> -version <version>
    swf kickoff -workflow-type <name,version> -workflow-id <id> -input <data> -queue <queue>
    swf workflow -list
    swf history

How can I get help?
    %s <command> -h

`, this.Cmd, this.Cmd)

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
