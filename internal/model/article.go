package model

import (
	"github.com/suisbuds/miao/pkg/app"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	TagID         int    `json:"tag_id"`
	Tag           Tag    `json:"tag"`
}

type ArticleRow struct {
	ArticleID          uint32
	TagID              uint32
	TagName            string
	ArticleTitle       string
	ArticleDescription string
	CoverImageUrl      string
	Content            string
}

func (a Article) TableName() string {
	return "mio_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}


