package swfapi

import (
	"fmt"
	"net/http"

	"github.com/funkygao/swf/models"
)

func (this *Client) RegisterWorkflowType(input *RegisterWorkflowTypeInput) (*RegisterWorkflowTypeOutput, error) {
	resp, body, err := this.call(models.OpRegisterWorkflowType, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &RegisterWorkflowTypeOutput{}
	output.from(body)
	return output, nil
}

func (this *Client) RegisterActivityType(input *RegisterActivityTypeInput) (*RegisterActivityTypeOutput, error) {
	resp, body, err := this.call(models.OpRegisterActivityType, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &RegisterActivityTypeOutput{}
	output.from(body)
	return output, nil
}

func (this *Client) StartWorkflowExecution(input *StartWorkflowExecutionInput) (*StartWorkflowExecutionOutput, error) {
	resp, body, err := this.call(models.OpRegisterActivityType, input)
	if len(err) >= 1 {
		return nil, err[0]
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", http.StatusText(resp.StatusCode))
	}

	output := &StartWorkflowExecutionOutput{}
	output.from(body)
	return output, nil
}

func (this *Client) PollForActivityTask() {

}

func (this *Client) PollForDecisionTask() {

}

func (this *Client) RespondActivityTaskCompleted() {

}

func (this *Client) RespondDecisionTaskCompleted() {

}

func (this *Client) ListWorkflowTypes() {

}

func (this *Client) ListWorkflowExecutions(openOrClose bool) {

}

func (this *Client) GetWorkflowExecutionHistory() {

}

func (this *Client) ListActivityTypes() {

}
