package natsclient

const (
	TopicTaskRcPrefix string = "task.Rc."

	TopicContentSuffix           string = "content.newId"
	TopicTradeSuffix             string = "trades.newId"
	TopicToProcessContentSuffix  string = "content.process.newId"
	TopicToRequestDecisionSuffix string = "decision.newId"
	TopicToSyncGraphSuffix       string = "graph.sync.newId"
	TopicReportSnapshotSuffix    string = "report.snapshot.newId"

	TopicContentNew           = TopicTaskRcPrefix + TopicContentSuffix
	TopicTradeNew             = TopicTaskRcPrefix + TopicTradeSuffix
	TopicContentToProcessNew  = TopicTaskRcPrefix + TopicToProcessContentSuffix
	TopicToRequestDecisionNew = TopicTaskRcPrefix + TopicToRequestDecisionSuffix
	TopicToSyncGraphNew       = TopicTaskRcPrefix + TopicToSyncGraphSuffix
	TopicReportSnapshot       = TopicTaskRcPrefix + TopicReportSnapshotSuffix
)
