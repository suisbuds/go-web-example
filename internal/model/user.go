package model



type User struct {
	*Model
	Username  string `json:"username"`
    Password  string `json:"password"`
    Avatar    string `json:"avatar"`
    UserType  uint8  `json:"user_type"`
    State     uint8  `json:"state"`
}

