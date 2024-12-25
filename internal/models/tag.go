package models

import "github.com/suisbuds/miao/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "miao_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}