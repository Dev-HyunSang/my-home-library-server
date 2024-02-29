package dto

import (
	"time"

	"github.com/dev-hyunsang/my-home-library-server/ent"
)

type RequestUpdateBook struct {
	Title             string `json:"title"`
	SubTitle          string `json:"sub_title"`
	Publisher         string `json:"publisher"`
	PublishingCompany string `json:"publishing_company"`
	Memo              string `json:"memo"`
	TotalPage         int    `json:"total_page"`
	CurrentPage       int    `json:"current_page"`
}

type ResponseUpdateBook struct {
	Status      Status    `json:"status"`
	Data        *ent.Book `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}

type ResponseGetBooks struct {
	Status      Status      `json:"status"`
	Data        []*ent.Book `json:"data"`
	RespondedAt time.Time   `json:"responded_at"`
}

type ResponseGetBook struct {
	Status      Status    `json:"status"`
	Data        *ent.Book `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}

type ResponseBook struct {
	Status      Status    `json:"status"`
	RespondedAt time.Time `json:"responded_at"`
}
