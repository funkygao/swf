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

type Decider struct {
	Ui  cli.Ui
	Cmd string

	cli *swfapi.Client
}

func (this *Decider) Run(args []string) (exitCode int) {
	cmdFlags := flag.NewFlagSet("decider", flag.ContinueOnError)
	cmdFlags.Usage = func() { this.Ui.Output(this.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	this.cli = swfapi.Default()
	this.mainLoop()

	return
}

func (this *Decider) mainLoop() {
	this.Ui.Info("enter decider main loop")
	var (
		pollInput = &models.PollForDecisionTaskInput{}
		decision  = &models.RespondDecisionTaskCompletedInput{}
	)
	pollInput.WorkflowType.Name = "w1"
	pollInput.WorkflowType.Version = "v1"
	for {
		pollOutput, err := this.cli.PollForDecisionTask(pollInput)
		if err != nil {
			this.Ui.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}

		this.Ui.Output(fmt.Sprintf("task token: %s %+v", pollOutput.TaskToken, pollOutput.Events))

		displayNotify("decider", pollOutput.TaskToken, nil)

		// worker orchestration according to history events
		decision.Reset()
		d := models.NewDecision(models.DecisionTypeScheduleActivityTask)
		d.ScheduleActivityTaskDecisionAttributes = &models.ScheduleActivityTaskDecisionAttributes{
			ActivityType: models.ActivityType{
				Name:    "sms",
				Version: "v1",
			},
			Input: "hello from decider",
		}

		decision.AddDecision(*d)

		// respond
		decision.TaskToken = pollOutput.TaskToken
		_, err = this.cli.RespondDecisionTaskCompleted(decision)
		if err != nil {
			this.Ui.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}

		time.Sleep(time.Second)
	}
}

func (*Decider) Synopsis() string {
	return fmt.Sprintf("(%s) Start the decider process.", color.Yellow("decider"))
}

func (this *Decider) Help() string {
	help := fmt.Sprintf(`
Usage: %s decider [options]

    %s   

`, this.Cmd, this.Synopsis())
	return strings.TrimSpace(help)
}
