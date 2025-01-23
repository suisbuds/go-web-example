package model

import (
	"github.com/suisbuds/miao/pkg/app"
)

type Article struct {
	*Model
	TagID         int    `json:"tag_id"`
	Tag           Tag    `json:"tag"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
}


func (a Article) TableName() string {
	return "mio_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}


