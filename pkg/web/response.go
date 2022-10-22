package web

//Reponse represents a general response
type Response struct {
	StatusCode   int         `json:"status_code"`
	Message      string      `json:"message"`
	Response     interface{} `json:"response,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}
