package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/golib/color"
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
	return fmt.Sprintf("(%s) Mark this alarm <being handled | handled | false alarm | phone called>.", color.Blue("worker"))
}

func (this *Marker) Help() string {
	help := fmt.Sprintf(`
Usage: %s marker [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
