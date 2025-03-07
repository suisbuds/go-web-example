package models

import (
	"github.com/suisbuds/miao/pkg/app"
	"gorm.io/gorm"
)

// 封装 Tag 模块对应的数据操作, 与 Dao 层交互, 借助 Gorm 操作数据库
// 显式错误处理, 让代码更具可读性

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// 实现 Gorm 的接口，指定 Tag 对应的 pg 表名
func (t Tag) TableName() string {
	return "miao_tag"
}

// Swagger 文档生成
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}


func (t Tag) Create(db *gorm.DB) error {
	if err := db.Create(&t).Error; err != nil {
		return err
	}
	return nil
}

// 根据 ID 获取 Tag
func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	if err := db.Model(&t).Where("id = ? AND is_del = ?", t.Model.ID, 0).Update("is_del", 1).Error; err != nil {
		return err
	}
	if err := db.Where("id = ?", t.Model.ID).Delete(&t).Error; err != nil {
		return err
	}
	return nil
}

// 统计符合条件的 Tag 数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	// Count 参数传入 int64, 但是 app.go 的参数 TotalRows 是 int, 需要显式类型转换
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// 获取 Tag 列表
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	// where 筛选过滤
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// 根据 ID 列表获取对应 Tag 列表
func (t Tag) ListByIDs(db *gorm.DB, ids []uint32) ([]*Tag, error) {
	var tags []*Tag
	db = db.Where("state = ? AND is_del = ?", t.State, 0)
	err := db.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}