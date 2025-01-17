package model

import "gorm.io/gorm"

type ArticleTag struct {
	*Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "miao_article_tag"
}

func (a ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a ArticleTag) GetByArticleID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.Model.ID, 0).Update("is_del", 1).Error; err != nil {
		return err
	}
	if err := db.Where("id = ?", a.Model.ID).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	if err := db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Update("is_del", 1).Error; err != nil {
		return err
	}
	if err := db.Where("article_id = ?", a.ArticleID).Delete(&a).Limit(1).Error; err != nil {
		return err
	}

	return nil
}

func (a ArticleTag) ListByTagID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

// 根据文章ID列表获取文章标签列表
func (a ArticleTag) ListByArticleIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleTags, nil
}
