package dto

import "time"

type ErrStatus struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	ErrMessage string `json:"err_message"`
}

type ResponseErr struct {
	Status      ErrStatus `json:"status"`
	RespondedAt time.Time `json:"responded_at"`
}
