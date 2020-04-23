package kinesis_consumer

type KinesisAction string

const (
	KinesisActionCreate KinesisAction = "create"
	KinesisActionUpdate               = "update"
	KinesisActionDelete               = "delete"
	KinesisActionNotify               = "notify"
)
