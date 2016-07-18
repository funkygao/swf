package models

import (
	"time"
)

type ActivityType struct {
	Name        string `db:"name"`
	Version     string `db:"ver"`
	Domain      string `db:"domain"`
	Description string `db:"desc"`
	Cluster     string `db:"cluster"`

	DefaultTaskHeartbeatTimeout       time.Duration `db:"-"`
	DefaultTaskScheduleToCloseTimeout time.Duration `db:"-"`
	DefaultTaskScheduleToStartTimeout time.Duration `db:"-"`
	DefaultTaskStartToCloseTimeout    time.Duration `db:"-"`
}

func (this ActivityType) Topic() string {
	return activityTopicPrefix + this.Name
}

func (this ActivityType) Validate() error {
	if this.Name == "" || this.Version == "" || this.Cluster == "" {
		return ErrRequiredMissing
	}

	return nil
}
