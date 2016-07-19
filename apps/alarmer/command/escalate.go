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

type Escalate struct {
	Ui  cli.Ui
	Cmd string

	cli *swfapi.Client
}

func (this *Escalate) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("escalate", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	this.cli = swfapi.Default()
	this.mainLoop()

	return
}

func (this *Escalate) mainLoop() {
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

func (*Escalate) Synopsis() string {
	return fmt.Sprintf("(%s) Escalate this alarm to a higher level.", color.Blue("worker"))
}

func (this *Escalate) Help() string {
	help := fmt.Sprintf(`
Usage: %s escalate [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
