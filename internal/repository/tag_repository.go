package repository

import (
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/pkg/app"
)

func (d *Repository) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name: name,
		Model: &model.Model{
			CreatedBy: createdBy,
			State:     state,
		},
	}

	return tag.Create(d.engine)
}

func (d *Repository) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id, State: state}}
	return tag.Get(d.engine)
}

func (d *Repository) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
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

func (d *Repository) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

func (d *Repository) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, Model: &model.Model{State: state}}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Repository) GetTagListByIDs(ids []uint32, state uint8) ([]*model.Tag, error) {
	tag := model.Tag{Model: &model.Model{State: state}}
	return tag.ListByIDs(d.engine, ids)
}

func (d *Repository) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, Model: &model.Model{State: state}}
	return tag.Count(d.engine)
}
