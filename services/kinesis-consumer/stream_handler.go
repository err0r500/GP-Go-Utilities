package kinesis_consumer

type StreamHandler interface {
	HandleStreamRecord(action KinesisAction, record interface{})
}
