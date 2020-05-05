package actiontype

// Argument 環境変数を管理する
type Argument struct {
	slackWebhookURL5min string
}

// NewArgument Argumentを作る
func NewArgument(slackWebhookURL5min string) Argument {
	return Argument{slackWebhookURL5min: slackWebhookURL5min}
}

// GetSlackWebhookURL5min SlackWebhookURL5min
func (a Argument) GetSlackWebhookURL5min() string {
	return a.slackWebhookURL5min
}
