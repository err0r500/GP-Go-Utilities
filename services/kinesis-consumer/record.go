package kinesis_consumer

type Record struct {
	From   string                 `json:"from"`
	Action KinesisAction          `json:"action"`
	Data   map[string]interface{} `json:"data"`
}
