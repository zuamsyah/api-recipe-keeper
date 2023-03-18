package models

import (
	"time"

	"api-recipe-keeper/helpers"
)

type Ingridient struct {
	ID        int       `json:"id,omitempty" gorm:"AUTO_INCREMENT" example:"115"`
	Name      string    `json:"name,omitempty" gorm:"type:varchar(255)" validate:"required,max=255"`
	Unit      string    `json:"unit,omitempty" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}

type IngridientShort struct {
	ID       int    `json:"id"`
	Quantity string `json:"quantity"`
}

func (o *Ingridient) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "ingridients", "as": "i"},
		"fields": map[string]map[string]string{
			"id":         {"name": "i.id", "as": "id"},
			"name":       {"name": "i.name", "as": "name"},
			"unit":       {"name": "i.unit", "as": "unit"},
			"created_at": {"name": "i.created_at", "as": "created_at"},
			"updated_at": {"name": "i.updated_at", "as": "updated_at"},
		},
	}
}

func (o *Ingridient) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "ingridients", "id", id, params, o.Schema(), map[string]interface{}{})
}

func (o *Ingridient) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *Ingridient) Create(ctx helpers.Context) map[string]interface{} {
	isValid, msg := helpers.Validate(ctx, o)
	if !isValid {
		return msg
	}
	params, err := o.SetDefaultValue(ctx)
	if err != nil {
		return params
	}
	helpers.GetDB(ctx).Create(o)
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *Ingridient) UpdateById(ctx helpers.Context) map[string]interface{} {
	helpers.GetDB(ctx).Model(Ingridient{}).Where("id = ?", o.ID).Updates(o)
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *Ingridient) DeleteById(ctx helpers.Context) map[string]interface{} {
	id := helpers.Convert(o.ID).String()
	helpers.GetDB(ctx).Model(Ingridient{}).Where("id = ?", o.ID).Delete(&Ingridient{})
	return helpers.DeletedMessage("cities", "id", id)
}

func (o *Ingridient) SetDefaultValue(ctx helpers.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	return params, nil
}
