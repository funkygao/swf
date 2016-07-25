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

		displayNotify("decider", fmt.Sprintf("workflow:%s-%s, task:%s", pollOutput.WorkflowType.Name,
			pollOutput.WorkflowType.Version, pollOutput.TaskToken), nil)

		// worker orchestration according to history events
		decision.Reset()
		decision.TaskToken = pollOutput.TaskToken
		this.decide(pollOutput.TaskToken, pollOutput.Events, decision)

		// respond
		_, err = this.cli.RespondDecisionTaskCompleted(decision)
		if err != nil {
			this.Ui.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}

		time.Sleep(time.Second)
	}
}

func (*Decider) decide(taskToken string, events models.HistoryEvents, decision *models.RespondDecisionTaskCompletedInput) {
	d := models.NewDecision(models.DecisionTypeScheduleActivityTask)
	d.ScheduleActivityTaskDecisionAttributes = &models.ScheduleActivityTaskDecisionAttributes{
		ActivityType: deciderActivityType,
		Input:        "hello from decider",
	}

	decision.AddDecision(*d)

	d = models.NewDecision(models.DecisionTypeScheduleActivityTask)
	d.ScheduleActivityTaskDecisionAttributes = &models.ScheduleActivityTaskDecisionAttributes{
		ActivityType: markerActivityType,
		Input:        "hello1 from decider",
	}
	decision.AddDecision(*d)
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
