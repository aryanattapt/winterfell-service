package model

type ErrorResponse struct {
	Code       int      `json:"code"`       // HTTP status code indicating error or success
	Message    string   `json:"message"`    // A human-readable message explaining the error
	Stacktrace []string `json:"stacktrace"` // An empty slice (You can add actual stack trace if needed)
	Timestamp  string   `json:"timestamp"`  // Current timestamp in ISO 8601 format
	Path       string   `json:"path"`       // The API endpoint path where the error occurred
}

type SuccessResponse struct {
	SuccessResponseNoData
	Metadata interface{} `json:"metadata"`
	Data     interface{} `json:"data"`
}

type SuccessResponseNoData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
