package actiontype

type Argument struct {
	slackWebhookUrl5min string
}

func (a Argument) GetSlackWebhookUrl5min() string {
	return a.slackWebhookUrl5min
}
