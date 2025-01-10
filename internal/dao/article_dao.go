package dao

import (
	"github.com/suisbuds/miao/internal/models"
	"github.com/suisbuds/miao/pkg/app"
)

// 可以改为 Repository 模式，将数据操作封装到 Repository 中，Dao 层只负责调用 Repository

//  定义多个 Article 用于耦合 Article 和 Tag 以处理不同参数, 设计上的失败
type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Description          string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*models.Article, error) {
	article := models.Article{
		Title:         param.Title,
		Description:          param.Description,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &models.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) GetArticle(id uint32, state uint8) (models.Article, error) {
	article := models.Article{Model: &models.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := models.Article{Model: &models.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Description != "" {
		values["description"] = param.Description
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}


func (d *Dao) DeleteArticle(id uint32) error {
	article := models.Article{Model: &models.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*models.ArticleRow, error) {
	article := models.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}

func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int, error) {
	article := models.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

