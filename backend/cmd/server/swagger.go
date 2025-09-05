package main

import "time"

// @swagger:model
type DeletedAt struct {
	Time  time.Time `json:"time,omitempty"`
	Valid bool      `json:"valid,omitempty"`
}

// @swagger:model
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// These prevent "unused" warnings
var _ = DeletedAt{}
var _ = APIResponse{}
