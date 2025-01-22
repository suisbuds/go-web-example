package model


type Menu struct {
	*Model
	MenuName   string `json:"menu_name"`
    URL        string `json:"url"`
    ParentID   int `json:"parent_id"`
    ParentName string `json:"parent_name"`
    Level      int    `json:"level"`
}
