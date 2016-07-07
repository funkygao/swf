package memory

import (
	"testing"

	"github.com/funkygao/assert"
	"github.com/funkygao/swf/models"
)

func TestActivityTypeRelated(t *testing.T) {
	m := New()
	err := m.Start()
	assert.Equal(t, nil, err)
	err = m.RegisterActivityType(&models.ActivityType{
		Name:    "foo",
		Version: "bar",
	})
	assert.Equal(t, nil, err)

	t0 := &models.ActivityType{
		Name:    "notfound",
		Version: "invalid",
	}
	t1, err := m.GetActivityType(t0)
	assert.Equal(t, nil, err)

	t0 = &models.ActivityType{
		Name:    "foo",
		Version: "bar",
	}
	t1, err = m.GetActivityType(t0)
	assert.Equal(t, nil, err)
	assert.Equal(t, t0.Name, t1.Name)
	assert.Equal(t, t0.Version, t1.Version)

}

func TestWorkflowTypeRelated(t *testing.T) {
	m := New()
	err := m.Start()
	assert.Equal(t, nil, err)
	err = m.RegisterWorkflowType(&models.WorkflowType{
		Name:    "foo",
		Version: "bar",
	})
	assert.Equal(t, nil, err)

	t0 := &models.WorkflowType{
		Name:    "notfound",
		Version: "invalid",
	}
	t1, err := m.GetWorkflowType(t0)
	assert.Equal(t, nil, err)

	t0 = &models.WorkflowType{
		Name:    "foo",
		Version: "bar",
	}
	t1, err = m.GetWorkflowType(t0)
	assert.Equal(t, nil, err)
	assert.Equal(t, t0.Name, t1.Name)
	assert.Equal(t, t0.Version, t1.Version)

}
