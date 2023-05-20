package natsclient

const (
	TopicTaskRskcPrefix        = "task.rskc."
	TopicContentSuffix  string = "content.newId"
	TopicTradeSuffix    string = "trades.newId"
	TopicContentNew            = TopicTaskRskcPrefix + TopicContentSuffix
	TopicTradeNew              = TopicTaskRskcPrefix + TopicTradeSuffix
)
