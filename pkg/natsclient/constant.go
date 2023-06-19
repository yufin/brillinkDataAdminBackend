package natsclient

const (
	TopicTaskRskcPrefix string = "task.rskc."

	TopicContentSuffix           string = "content.newId"
	TopicTradeSuffix             string = "trades.newId"
	TopicToProcessContentSuffix  string = "content.process.newId"
	TopicToRequestDecisionSuffix string = "decision.newId"
	TopicToSyncGraphSuffix       string = "graph.sync.newId"

	TopicContentNew           = TopicTaskRskcPrefix + TopicContentSuffix
	TopicTradeNew             = TopicTaskRskcPrefix + TopicTradeSuffix
	TopicContentToProcessNew  = TopicTaskRskcPrefix + TopicToProcessContentSuffix
	TopicToRequestDecisionNew = TopicTaskRskcPrefix + TopicToRequestDecisionSuffix
	TopicToSyncGraphNew       = TopicTaskRskcPrefix + TopicToSyncGraphSuffix
)
