package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/swf-go/v1"
)

type Activity struct {
	Ui  cli.Ui
	Cmd string

	activityType models.ActivityType
}

func (this *Activity) Run(args []string) (exitCode int) {
	var (
		registerMode, listMode bool
	)
	cmdFlags := flag.NewFlagSet("activity", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	cmdFlags.BoolVar(&registerMode, "register", false, "")
	cmdFlags.BoolVar(&listMode, "list", true, "")
	cmdFlags.StringVar(&this.activityType.Name, "name", "", "")
	cmdFlags.StringVar(&this.activityType.Version, "version", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	switch {
	case listMode:
		this.listActivityTypes()

	case registerMode:
		this.registerActivityType()
	}

	return
}

func (this *Activity) listActivityTypes() {
	swfapi.Default().ListActivityTypes()
}

func (this *Activity) registerActivityType() {
	input := &models.RegisterActivityTypeInput{
		ActivityType: this.activityType,
	}
	_, err := swfapi.Default().RegisterActivityType(input)
	if err != nil {
		this.Ui.Error(err.Error())
		return
	}

	this.Ui.Info("registered")
}

func (*Activity) Synopsis() string {
	return "Register/List/Modify activity and activity type."
}

func (this *Activity) Help() string {
	help := fmt.Sprintf(`
Usage: %s activity [options]

    %s

    -register <name>
      Registers  a new activity type.

      -version <version>

      [-description <value>]

    -list

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
