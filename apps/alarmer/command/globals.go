package command

import (
	"github.com/funkygao/swf/models"
)

var (
	deciderActivityType = models.ActivityType{
		Name:    "decider",
		Version: "v1",
	}

	markerActivityType = models.ActivityType{
		Name:    "marker",
		Version: "v1",
	}

	smsActivityType = models.ActivityType{
		Name:    "sms",
		Version: "v1",
	}
)
