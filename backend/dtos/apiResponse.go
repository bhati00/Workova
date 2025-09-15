package dtos

type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message,omitempty" example:"Operation successful"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"Something went wrong"`
}

// BatchResult represents the result of batch operations
type BatchResult struct {
	TotalProcessed int          `json:"total_processed"`
	Successful     int          `json:"successful"`
	Failed         int          `json:"failed"`
	Errors         []BatchError `json:"errors,omitempty"`
}
type BatchDeleteRequest struct {
	IDs []uint `json:"ids"`
}

// BatchError represents individual batch operation errors
type BatchError struct {
	Index int    `json:"index"`
	ID    uint   `json:"id,omitempty"`
	Error string `json:"error"`
}
