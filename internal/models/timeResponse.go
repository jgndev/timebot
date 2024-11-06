package models

type TimeResponse struct {
	UTC   string            `json:"utc"`
	Times map[string]string `json:"times"`
}
