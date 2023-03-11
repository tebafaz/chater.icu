package models

type StatusError struct {
	Error string `json:"error"`
}

type StatusSuccess struct {
	Success string `json:"success"`
}
