package natsclient

type Activity interface {
	OnClose() error
	Setup() error
}

func InitActivity(activity Activity) error {
	if err := activity.Setup(); err != nil {
		return err
	}
	return nil
}

func CloseActivity(activity Activity) error {
	if err := activity.OnClose(); err != nil {
		return err
	}
	return nil
}
