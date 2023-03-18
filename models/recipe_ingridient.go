package models

import (
	"time"

	"api-recipe-keeper/helpers"
)

type RecipeIngridient struct {
	ID           int       `json:"id,omitempty" gorm:"AUTO_INCREMENT" example:"1"`
	IDRecipe     int       `json:"-" gorm:"index:recipes_id_recipe"`
	IDIngridient int       `json:"-" gorm:"index:recipes_id_ingridient"`
	Quantity     string    `json:"name" gorm:"type:varchar(150)"`
	CreatedAt    time.Time `json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}

func (o *RecipeIngridient) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "couriers"},
		"fields": map[string]map[string]string{
			"id":                           {"name": "id", "as": "id", "type": "int"},
			"code":                         {"name": "code", "as": "code"},
			"name":                         {"name": "name", "as": "name"},
			"provider":                     {"name": "provider", "as": "provider"},
			"delivery_cost_estimation_key": {"name": "delivery_cost_estimation_key", "as": "delivery_cost_estimation_key"},
			"delivery_tracking_type":       {"name": "delivery_tracking_type", "as": "delivery_tracking_type"},
			"pickup_service":               {"name": "pickup_service", "as": "pickup_service"},
			"image_url":                    {"name": "image_url", "as": "image_url"},
			"sort_number":                  {"name": "sort_number", "as": "sort_number", "type": "int"},
			"is_active":                    {"name": "is_active", "as": "is_active", "type": "bool"},
		},
	}
}

func (o *RecipeIngridient) GetByCode(ctx helpers.Context, code string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "couriers", "code", code, params, o.Schema(), map[string]interface{}{})
}

func (o *RecipeIngridient) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "couriers", "id", id, params, o.Schema(), map[string]interface{}{})
}

func (o *RecipeIngridient) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *RecipeIngridient) Create(ctx helpers.Context) map[string]interface{} {
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

func (o *RecipeIngridient) UpdateById(ctx helpers.Context) map[string]interface{} {
	helpers.GetDB(ctx).Model(RecipeIngridient{}).Where("id = ?", o.ID).Updates(o)
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *RecipeIngridient) DeleteById(ctx helpers.Context) map[string]interface{} {
	id := helpers.Convert(o.ID).String()
	helpers.GetDB(ctx).Model(RecipeIngridient{}).Where("id = ?", o.ID).Delete(&RecipeIngridient{})
	return helpers.DeletedMessage("couriers", "id", id)
}

func (o *RecipeIngridient) SetDefaultValue(ctx helpers.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	return params, nil
}
