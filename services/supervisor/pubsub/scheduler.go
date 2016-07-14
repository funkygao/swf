package pubsub

func (this *Supervisor) schedule() {
	for {
		select {
		case <-this.quit:
			return

		case <-this.decisionCh:

		case <-this.notificationCh:
		}
	}

}
