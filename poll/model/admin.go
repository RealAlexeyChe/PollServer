package model

type AdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Id       string `json:"id"`
}
