package dtos

type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message,omitempty" example:"Operation successful"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"Something went wrong"`
}
type BatchDeleteRequest struct {
	IDs    []uint   `json:"ids"`
	JobIDs []string `json:"job_ids"`
}
