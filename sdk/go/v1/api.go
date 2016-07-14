package swfapi

import (
	"fmt"
	"net/http"

	"github.com/funkygao/swf/models"
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
	output.From(body)
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

func (this *Client) ListActivityTypes() {

}

func (this *Client) ListWorkflowExecutions(openOrClose bool) {

}

func (this *Client) GetWorkflowExecutionHistory() {

}
