package model


type RoleMenu struct {
    *Model
    RoleID int `json:"role_id"`
    MenuID int `json:"menu_id"`
}