package model


type RoleMenu struct {
    *Model
    RoleID uint32 `json:"role_id"`
    MenuID uint32 `json:"menu_id"`
}