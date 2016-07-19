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

type Phone struct {
	Ui  cli.Ui
	Cmd string

	cli *swfapi.Client
}

func (this *Phone) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("phone", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	this.cli = swfapi.Default()
	this.mainLoop()

	return
}

func (this *Phone) mainLoop() {
	this.Ui.Info("enter worker main loop")
	var (
		pollInput    = &models.PollForActivityTaskInput{}
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

		// respond
		respondInput.TaskToken = pollOutput.TaskToken
		respondOutput, err := this.cli.RespondActivityTaskCompleted(respondInput)
		if err != nil {
			this.Ui.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}
		this.Ui.Output(fmt.Sprintf("%+v", respondOutput))
	}
}

func (*Phone) Synopsis() string {
	return fmt.Sprintf("(%s) Manually phone call the person in charge.", color.Blue("worker"))
}

func (this *Phone) Help() string {
	help := fmt.Sprintf(`
Usage: %s phone [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
