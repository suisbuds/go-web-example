package service

import (
	"github.com/suisbuds/miao/internal/model"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// 文章 ID 主键自动创建, 这里传入的是 TagID
type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Description   string `form:"description" binding:"omitempty,min=2,max=255"`
	Content       string `form:"content" binding:"omitempty,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"omitempty,url"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=2,max=100"`
	Description   string `form:"description" binding:"omitempty,min=2,max=255"`
	Content       string `form:"content" binding:"omitempty,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"omitempty,url"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}
