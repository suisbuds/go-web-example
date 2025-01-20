package model


type Menu struct {
	*Model
	MenuName   string `json:"menu_name"`
    URL        string `json:"url"`
    ParentID   uint32 `json:"parent_id"`
    ParentName string `json:"parent_name"`
    Level      int    `json:"level"`
    State      uint8  `json:"state"`
}
