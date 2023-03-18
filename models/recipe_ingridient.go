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
		"table": map[string]string{"name": "recipe_ingridients", "as": "ri"},
		"fields": map[string]map[string]string{
			"id":            {"name": "ri.id", "as": "id", "type": "int"},
			"id_recipe":     {"name": "ri.id_recipe", "as": "id_recipe", "is_hide": "true"},
			"id_ingridient": {"name": "ri.id_ingridient", "as": "id_ingridient", "is_hide": "true"},
			"name":          {"name": "ri.name", "as": "name"},
			"unit":          {"name": "ri.unit", "as": "unit"},
		},
	}
}

func (o *RecipeIngridient) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "recipe_ingridients", "id", id, params, o.Schema(), map[string]interface{}{})
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
	ri := RecipeIngridient{}
	if helpers.GetDB(ctx).Model(RecipeIngridient{}).Where("id_recipe = ? AND id_ingridient = ?", o.IDRecipe, o.IDIngridient).First(&ri).RowsAffected == 0 {
		helpers.GetDB(ctx).Create(o)
	} else {
		helpers.GetDB(ctx).Model(RecipeIngridient{}).Where("id = ?", ri.ID).Updates(o)
	}

	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *RecipeIngridient) SetDefaultValue(ctx helpers.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	return params, nil
}
