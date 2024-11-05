package models

type ArticleTags struct {
	*Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
}

func (a ArticleTags) TableName() string {
	return "miao_article_tags"
}
