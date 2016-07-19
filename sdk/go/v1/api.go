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
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &models.RegisterWorkflowTypeOutput{}
	err = output.From(body)
	return output, err
}

func (this *Client) RegisterActivityType(input *models.RegisterActivityTypeInput) (*models.RegisterActivityTypeOutput, error) {
	resp, body, err := this.invoke(models.OpRegisterActivityType, input)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &models.RegisterActivityTypeOutput{}
	err = output.From(body)
	return output, err
}

func (this *Client) StartWorkflowExecution(input *models.StartWorkflowExecutionInput) (*models.StartWorkflowExecutionOutput, error) {
	resp, body, err := this.invoke(models.OpStartWorkflowExecution, input)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &models.StartWorkflowExecutionOutput{}
	err = output.From(body)
	return output, err
}

func (this *Client) PollForActivityTask(input *models.PollForActivityTaskInput) (output *models.PollForActivityTaskOutput, err error) {
	resp, body, err := this.invoke(models.OpPollForActivityTask, input)
	if err != nil {
		return nil, err
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

func (this *Client) PollForDecisionTask(input *models.PollForDecisionTaskInput) (output *models.PollForDecisionTaskOutput, err error) {
	resp, body, err := this.invoke(models.OpPollForDecisionTask, input)
	if err != nil {
		return nil, err
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
	if err != nil {
		return nil, err
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
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output = &models.RespondDecisionTaskCompletedOutput{}
	err = output.From(body)
	return
}

func (this *Client) invoke(op string, payload interface{}) (gorequest.Response, []byte, error) {
	agent := gorequest.New()
	resp, body, errs := agent.Post(this.cf.Endpoint()).
		Set("User-Agent", "swf-go:"+swf.Version).
		Set("X-Swf-Api", op).
		SendStruct(payload). // json encode
		EndBytes()
	var err error
	if len(errs) > 0 {
		err = errs[0]
	}
	return resp, body, err
}
