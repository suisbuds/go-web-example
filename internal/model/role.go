package model


type Role struct {
	*Model
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Value    string `json:"value"`
}