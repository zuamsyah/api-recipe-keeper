package models

import (
	"time"

	"api-recipe-keeper/helpers"
)

type FoodCategory struct {
	ID        int       `json:"id,omitempty" gorm:"AUTO_INCREMENT" example:"115"`
	Name      string    `json:"name,omitempty" gorm:"type:varchar(255)" validate:"required,max=255" example:"Kota Depok"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}

func (o *FoodCategory) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "food_categories", "as": "fc"},
		"fields": map[string]map[string]string{
			"id":         {"name": "fc.id", "as": "id"},
			"name":       {"name": "fc.name", "as": "name"},
			"created_at": {"name": "fc.created_at", "as": "created_at"},
			"updated_at": {"name": "fc.updated_at", "as": "updated_at"},
		},
	}
}

func (o *FoodCategory) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "food_categories", "id", id, params, o.Schema(), map[string]interface{}{})
}

func (o *FoodCategory) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *FoodCategory) Create(ctx helpers.Context) map[string]interface{} {
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

func (o *FoodCategory) UpdateById(ctx helpers.Context) map[string]interface{} {
	helpers.GetDB(ctx).Model(FoodCategory{}).Where("id = ?", o.ID).Updates(o)
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *FoodCategory) DeleteById(ctx helpers.Context) map[string]interface{} {
	id := helpers.Convert(o.ID).String()
	helpers.GetDB(ctx).Model(FoodCategory{}).Where("id = ?", o.ID).Delete(&FoodCategory{})
	return helpers.DeletedMessage("cities", "id", id)
}

func (o *FoodCategory) SetDefaultValue(ctx helpers.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	return params, nil
}
