package models

type ArticleTag struct {
	*Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "miao_article_tag"
}
