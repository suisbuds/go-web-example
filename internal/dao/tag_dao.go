package dao

import (
	"github.com/suisbuds/miao/internal/models"
	"github.com/suisbuds/miao/pkg/app"
)


func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := models.Tag{
		Name:  name,
		State: state,
		Model: &models.Model{
			CreatedBy: createdBy,
		},
	}

	return tag.Create(d.engine)
}

func (d *Dao) GetTag(id uint32, state uint8) (models.Tag, error) {
	tag := models.Tag{Model: &models.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}


func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := models.Tag{
		Model: &models.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := models.Tag{Model: &models.Model{ID: id}}
	return tag.Delete(d.engine)
}


func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*models.Tag, error) {
	tag := models.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetTagListByIDs(ids []uint32, state uint8) ([]*models.Tag, error) {
	tag := models.Tag{State: state}
	return tag.ListByIDs(d.engine, ids)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := models.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}