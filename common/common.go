package common

// GokiRequest is for all requests
type GokiRequest struct {
	Command string `json:"command"`
	Args []string `json:"args"`
}

// GokiResponse is for all responses
type GokiResponse struct {
	Result string `json:"result"`
	TimeStamp string `json:"timestamp"`
}
