package command

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/funkygao/gocli"
	"github.com/funkygao/golib/color"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/sdk/go/v1"
)

type Marker struct {
	Ui  cli.Ui
	Cmd string

	cli *swfapi.Client
}

func (this *Marker) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("marker", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	this.cli = swfapi.Default()
	this.mainLoop()

	return
}

func (this *Marker) mainLoop() {
	this.Ui.Info("enter worker main loop")
	var (
		pollInput = &models.PollForActivityTaskInput{
			ActivityType: markerActivityType,
		}
		respondInput = &models.RespondActivityTaskCompletedInput{}
	)

	for {
		pollOutput, err := this.cli.PollForActivityTask(pollInput)
		if err != nil {
			this.Ui.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}

		this.Ui.Output(fmt.Sprintf("task token: %s (%s)", pollOutput.TaskToken, pollOutput.Input))

		// work hard
		displayNotify("marker", pollOutput.TaskToken, nil)

		state, _ := this.Ui.Ask(fmt.Sprintf("token:%s state:<handled|ignore>", pollOutput.TaskToken))
		respondInput.Result = state

		// respond
		respondInput.TaskToken = pollOutput.TaskToken
		_, err = this.cli.RespondActivityTaskCompleted(respondInput)
		if err != nil {
			this.Ui.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}
	}
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
