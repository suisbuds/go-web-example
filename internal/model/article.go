package model

import (
	"github.com/suisbuds/miao/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	TagID         int    `json:"tag_id"`
	Tag           Tag    `json:"tag"`
	State         uint8  `json:"state"`
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
	return "miao_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	err := db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}

	return article, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.Model.ID, 0).Update("is_del", 1).Error; err != nil {
		return err
	}
	if err := db.Where("id = ?", a.Model.ID).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

// 根据 TagID 获取文章列表
// 利用 Gorm 调用 Pg/MySql 注意两者 SQL 语法差异
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {

	// 定义查询字段
	fields := []string{
		"ar.id AS article_id",
		"ar.title AS article_title",
		"ar.description AS article_description",
		"ar.cover_image_url",
		"ar.content",
		"t.id AS tag_id",
		"t.name AS tag_name",
	}

	// 设置分页
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	articleTagTable := ArticleTag{}.TableName()
	tagTable := Tag{}.TableName()
	articleTable := Article{}.TableName()

	// 指定查询表和查询字段, 并将 Tag 表和 Article_Tag 表关联, Article 表和 Article_Tag 表关联, 然后 WHERE 筛选过滤
	// SQL 查询语法
	rows, err := db.Select(fields).
		Table(articleTagTable+" AS at").
		Joins("LEFT JOIN "+tagTable+" AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN "+articleTable+" AS ar ON at.article_id = ar.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()

	if err != nil {
		return nil, err
	}
	// 及时释放数据库资源
	defer rows.Close()

	// 遍历结果集, 构建返回数据
	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(
			&r.ArticleID,
			&r.ArticleTitle,
			&r.ArticleDescription,
			&r.CoverImageUrl,
			&r.Content,
			&r.TagID,
			&r.TagName,
		); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}

	return articles, nil
}

// 根据 TagID 统计文章数量
func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int64

	// 定义表名和别名
	articleTagTable := ArticleTag{}.TableName()
	tagTable := Tag{}.TableName()
	articleTable := Article{}.TableName()

	// 数据库查询并统计数量 (链式调用)
	err := db.Table(articleTagTable+" AS at").
		Joins("LEFT JOIN "+tagTable+" AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN "+articleTable+" AS ar ON at.article_id = ar.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
