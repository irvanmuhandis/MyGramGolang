package response

type ErrorMessage struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
