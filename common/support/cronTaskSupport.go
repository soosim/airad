package support

type CronTask struct {
	Name     string
	Spec     string
	WorkFunc func()
}
