package natsclient

const (
	TopicTaskRskcPrefix                = "task.rskc."
	TopicContentSuffix          string = "content.newId"
	TopicTradeSuffix            string = "trades.newId"
	TopicToProcessContentSuffix string = "content.process.newId"
	TopicContentNew                    = TopicTaskRskcPrefix + TopicContentSuffix
	TopicTradeNew                      = TopicTaskRskcPrefix + TopicTradeSuffix
	TopicContentToProcessNew           = TopicTaskRskcPrefix + TopicToProcessContentSuffix
)
