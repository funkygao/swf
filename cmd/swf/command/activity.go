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

type Activity struct {
	Ui  cli.Ui
	Cmd string

	listMode            bool
	zone, cluster       string
	regName, regVersion string
}

func (this *Activity) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("activity", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	cmdFlags.StringVar(&this.regName, "register", "", "")
	cmdFlags.StringVar(&this.regVersion, "version", "v1", "")
	cmdFlags.BoolVar(&this.listMode, "list", false, "")
	cmdFlags.StringVar(&this.zone, "z", ctx.DefaultZone(), "")
	cmdFlags.StringVar(&this.cluster, "c", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if this.listMode {
		this.listActivityTypes()
		return
	}

	if this.regName != "" {
		this.registerActivityType()
		return
	}

	return
}

func (this *Activity) listActivityTypes() {
	swfapi.WithZone(this.zone).ListActivityTypes()
}

func (this *Activity) registerActivityType() {
	input := &models.RegisterActivityTypeInput{}
	input.Cluster = this.cluster
	input.Name = this.regName
	input.Version = this.regVersion

	_, err := swfapi.WithZone(this.zone).RegisterActivityType(input)
	if err != nil {
		this.Ui.Error(err.Error())
		return
	}

	this.Ui.Info("activity type registered")
}

func (*Activity) Synopsis() string {
	return "Register/List/Modify activity type."
}

func (this *Activity) Help() string {
	help := fmt.Sprintf(`
Usage: %s activity -z <zone> -c <cluster> [options]

    %s

    -register <name>
      Registers  a new activity type.

      -version <version>


    -list

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
