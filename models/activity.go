package models

type ActivityType struct {
	Name    string
	Version string
	Cluster string
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
