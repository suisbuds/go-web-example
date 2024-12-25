package models

import "github.com/suisbuds/miao/pkg/app"

type Article struct {
    *Model
    Title         string `json:"title"`
    Description   string `json:"description"`
    Content       string `json:"content"`
    CoverImageUrl string `json:"cover_image_url"`
    State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "miao_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}