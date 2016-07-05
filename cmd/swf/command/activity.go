package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Activity struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Activity) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("activity", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Activity) Synopsis() string {
	return "Register/List/Manipulate activity and activity type."
}

func (this *Activity) Help() string {
	help := fmt.Sprintf(`
Usage: %s activity [options]

    %s

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
