package model


type Role struct {
	*Model
	UserID   uint32    `json:"user_id"`
	UserName string `json:"user_name"`
	Value    string `json:"value"`
}

func (r Role) TableName() string {
	return "mio_role"
}