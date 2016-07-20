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

type Sms struct {
	Ui  cli.Ui
	Cmd string

	cli *swfapi.Client
}

func (this *Sms) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("sms", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	this.cli = swfapi.Default()
	this.mainLoop()

	return
}

func (this *Sms) mainLoop() {
	this.Ui.Info("enter worker main loop")
	var (
		pollInput    = &models.PollForActivityTaskInput{}
		respondInput = &models.RespondActivityTaskCompletedInput{}
	)

	pollInput.ActivityType.Name = "sms"
	pollInput.ActivityType.Version = "v1"
	for {
		time.Sleep(time.Second)

		pollOutput, err := this.cli.PollForActivityTask(pollInput)
		if err != nil {
			this.Ui.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}

		this.Ui.Output(fmt.Sprintf("task token: %s (%s)", pollOutput.TaskToken, pollOutput.Input))

		// work hard
		displayNotify("sms", pollOutput.TaskToken, nil)

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

func (*Sms) Synopsis() string {
	return fmt.Sprintf("(%s) Send sms to person in charge of the alarm.", color.Blue("worker"))
}

func (this *Sms) Help() string {
	help := fmt.Sprintf(`
Usage: %s sms [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
