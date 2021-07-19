package models

type ErrorResponse struct {
	ErrorMsg  string `json:"error_msg"`
	SubMsg    string `json:"sub_msg"`
	SubCode   string `json:"sub_code,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}
