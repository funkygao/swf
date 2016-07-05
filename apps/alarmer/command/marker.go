package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Marker struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Marker) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("marker", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return
}

func (*Marker) Synopsis() string {
	return "(worker) Mark this alarm <being handled | handled | false alarm | phone called>."
}

func (this *Marker) Help() string {
	help := fmt.Sprintf(`
Usage: %s marker [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
