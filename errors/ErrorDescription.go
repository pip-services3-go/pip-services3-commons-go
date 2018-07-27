package errors

type ErrorDescription struct {
	Type          string                 `json:"type"`
	Category      string                 `json:"category"`
	Status        int                    `json:"status"`
	Code          string                 `json:"code"`
	Message       string                 `json:"message"`
	Details       map[string]interface{} `json:"details"`
	CorrelationId string                 `json:"correlation_id"`
	Cause         string                 `json:"cause"`
	StackTrace    string                 `json:"stack_trace"`
}
