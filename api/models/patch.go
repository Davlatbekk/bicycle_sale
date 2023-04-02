package models

type PatchRequest struct {
	ID     int `json:"id"`
	Fields map[string]interface{}
}
