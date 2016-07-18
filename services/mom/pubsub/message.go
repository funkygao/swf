package pubsub

func (this *PubSub) Pub(appid, topic, ver string, msg []byte) error {
	return nil

}

func (this *PubSub) AddTopic(cluster, appid, topic, ver string) error {
	return nil
}

func (this *PubSub) Sub(appid, topic, ver string) {

}
