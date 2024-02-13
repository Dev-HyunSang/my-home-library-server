package dto

import "time"

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

type ResponseLoginUser struct {
	Status      Status    `json:"status"`
	Data        JwtData   `json:"data"`
	ResponsedAt time.Time `json:"repsonsed_at"`
}
type JwtData struct {
	AccessUUID   string `json:"access_uuid"`
	AccessToken  string `json:"access_token"`
	RefreshUUID  string `json:"refresh_uuid"`
	RefreshToken string `json:"refresh_token"`
}
