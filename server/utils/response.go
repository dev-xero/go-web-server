package utils

type Response struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Payload any    `json:"payload"`
}
