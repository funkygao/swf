package channel

func (this *Channel) AddTopic(cluster, appid, topic, ver string) error {
	return nil
}

func (this *Channel) Pub(appid, topic, ver string, msg []byte) error {
	this.chs[appid+topic+ver] <- msg
	return nil
}

func (this *Channel) Sub(appid, topic, ver string) []byte {
	return <-this.chs[appid+topic+ver]
}
