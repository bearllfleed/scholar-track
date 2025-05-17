package config

type FailedDeletion struct {
	MaxRetry  int `json:"maxRetry"`
	Interval  int `json:"interval"`
	BatchSize int `json:"batchSize"`
}
