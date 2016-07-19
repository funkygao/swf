package swfapi

import (
	"fmt"
	"net/http"

	"github.com/funkygao/swf"
	"github.com/funkygao/swf/models"
	"github.com/parnurzeal/gorequest"
)

func (this *Client) RegisterWorkflowType(input *models.RegisterWorkflowTypeInput) (*models.RegisterWorkflowTypeOutput, error) {
	resp, body, err := this.invoke(models.OpRegisterWorkflowType, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &models.RegisterWorkflowTypeOutput{}
	output.From(body)
	return output, nil
}

func (this *Client) RegisterActivityType(input *models.RegisterActivityTypeInput) (*models.RegisterActivityTypeOutput, error) {
	resp, body, err := this.invoke(models.OpRegisterActivityType, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &models.RegisterActivityTypeOutput{}
	output.From(body)
	return output, nil
}

func (this *Client) StartWorkflowExecution(input *models.StartWorkflowExecutionInput) (*models.StartWorkflowExecutionOutput, error) {
	resp, body, err := this.invoke(models.OpStartWorkflowExecution, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &models.StartWorkflowExecutionOutput{}
	err = output.From(body)
	return output, err
}

func (this *Client) PollForActivityTask(input *models.PollForActivityTaskInput) (output chan *models.PollForActivityTaskOutput, err error) {
	resp, body, err := this.invoke(models.OpPollForActivityTask, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output = &models.PollForActivityTaskOutput{}
	err = output.From(body)
	return
}

func (this *Client) PollForDecisionTask(input *models.PollForDecisionTaskInput) (output chan *models.PollForDecisionTaskOutput, err error) {
	resp, body, err := this.invoke(models.OpPollForDecisionTask, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output = &models.PollForDecisionTaskOutput{}
	err = output.From(body)
	return
}

func (this *Client) RespondActivityTaskCompleted(input *models.RespondActivityTaskCompletedInput) (output *models.RespondActivityTaskCompletedOutput, err error) {
	resp, body, err := this.invoke(models.OpRespondActivityTaskCompleted, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output = &models.RespondActivityTaskCompletedOutput{}
	err = output.From(body)
	return
}

func (this *Client) RespondDecisionTaskCompleted(input *models.RespondDecisionTaskCompletedInput) (output *models.RespondDecisionTaskCompletedOutput, err error) {
	resp, body, err := this.invoke(models.OpRespondDecisionTaskCompleted, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output = &models.RespondDecisionTaskCompletedOutput{}
	err = output.From(body)
	return
}

func (this *Client) invoke(op string, payload interface{}) (gorequest.Response, []byte, []error) {
	agent := gorequest.New()
	return agent.Post(this.cf.Endpoint()).
		Set("User-Agent", "swf-go:"+swf.Version).
		Set("X-Swf-Api", op).
		SendStruct(payload).
		EndBytes()
}
