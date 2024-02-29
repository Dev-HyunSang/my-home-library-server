package dto

import (
	"time"

	"github.com/dev-hyunsang/my-home-library-server/ent"
)

type RequestJoinUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

type RequestLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ==== Respones ====
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseJoinUser struct {
	Status      Status    `json:"status"`
	Data        *ent.User `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}

type ResponseLoginUser struct {
	Status      Status    `json:"status"`
	Data        JwtData   `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}

type JwtData struct {
	AccessUUID   string `json:"access_uuid"`
	AccessToken  string `json:"access_token"`
	RefreshUUID  string `json:"refresh_uuid"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseUserData struct {
	Status      Status    `json:"status"`
	Data        *ent.User `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}
