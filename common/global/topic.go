package global

type TopicType string

const (
	LoginLog    TopicType = "login_log_queue"
	RequestLog  TopicType = "request_log_queue"
	OperateLog  TopicType = "operate_log_queue"
	AbnormalLog TopicType = "abnormal_log_queue"
	ErrorLog    TopicType = "error_log_queue"
	ApiCheck    TopicType = "api_check_queue"
	Notice      TopicType = "notice_queue"
)
