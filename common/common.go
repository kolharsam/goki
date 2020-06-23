package common

import "time"

// GokiRequest is for all requests
type GokiRequest struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

// GokiResult will be used to help format the response
type GokiResult struct {
	Value    string `json:"value"`
	ToFormat bool   `json:"toFormat"`
}

// GokiResponse is for all responses
type GokiResponse struct {
	Result    GokiResult `json:"result"`
	TimeStamp time.Time  `json:"timestamp"`
}
