package models

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
