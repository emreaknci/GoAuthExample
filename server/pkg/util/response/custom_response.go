package response

type CustomResponse[T any] struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       T           `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}
