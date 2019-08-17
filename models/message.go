package models

// Message para el cliente
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
