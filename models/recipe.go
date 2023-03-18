package models

import (
	"time"

	"api-recipe-keeper/helpers"
)

type Recipe struct {
	ID           int               `json:"id,omitempty" gorm:"AUTO_INCREMENT" example:"1"`
	IDCategory   int               `json:"-" gorm:"index:recipes_id_category"`
	FoodCategory FoodCategoryShort `json:"food_category" gorm:"-"`
	Ingridients  []IngridientShort `json:"ingridients" gorm:"-"`
	Name         string            `json:"name" gorm:"type:varchar(100)" example:"Jalur Nugraha Ekakurir (JNE)"`
	Quantity     string            `json:"quantity" gorm:"-"`
	CreatedAt    time.Time         `json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt    time.Time         `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}

func (o *Recipe) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "recipes", "as": "r"},
		"fields": map[string]map[string]string{
			"id":            {"name": "r.id", "as": "id"},
			"category.id":   {"name": "r.id_category", "as": "id_category"},
			"category.name": {"name": "fc.name", "as": "category_name"},
			"name":          {"name": "r.name", "as": "name"},
		},
		"relations": []map[string]string{
			{"name": "food_categories", "as": "fc", "on": "fc.id = r.id_category", "type": "BelongsTo"},
		},
		"has_many_relations": map[string]map[string]interface{}{
			"recipe_ingridients": {
				"primary_key": "id",
				"foreign_key": "id_recipe",
				"schema": map[string]interface{}{
					"table": map[string]string{"name": "recipe_ingridients", "as": "ri"},
					"fields": map[string]map[string]string{
						"id":              {"name": "ri.id", "as": "id"},
						"id_recipe":       {"name": "ri.id_recipe", "as": "id_recipe", "is_hide": "true"},
						"id_ingridient":   {"name": "ri.id_ingridient", "as": "id_ingridient", "is_hide": "true"},
						"ingridient.id":   {"name": "rii.id", "as": "ingridient_id"},
						"ingridient.name": {"name": "rii.name", "as": "ingridient_name"},
						"ingridient.unit": {"name": "rii.unit", "as": "ingridient_unit"},
						"quantity":        {"name": "ri.quantity", "as": "quantity"},
					},
					"relations": []map[string]string{
						{"name": "ingridients", "as": "rii", "on": "rii.id = ri.id_ingridient", "type": "BelongsTo"},
					},
				},
			},
		},
	}
}

func (o *Recipe) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "recipes", "id", id, params, o.Schema(), map[string]interface{}{})
}

func (o *Recipe) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	params["is_include_has_many"] = []string{"true"}
	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *Recipe) Create(ctx helpers.Context) map[string]interface{} {
	isValid, msg := helpers.Validate(ctx, o)
	if !isValid {
		return msg
	}
	params, err := o.SetDefaultValue(ctx)
	if err != nil {
		return params
	}
	o.IDCategory = o.FoodCategory.ID
	helpers.GetDB(ctx).Create(o)
	for _, ingridient := range o.Ingridients {
		recipeIngridient := RecipeIngridient{}
		recipeIngridient.IDRecipe = o.ID
		recipeIngridient.IDIngridient = ingridient.ID
		recipeIngridient.Quantity = ingridient.Quantity
		recipeIngridient.Create(ctx)
	}
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *Recipe) UpdateById(ctx helpers.Context) map[string]interface{} {
	helpers.GetDB(ctx).Model(Recipe{}).Where("id = ?", o.ID).Updates(o)
	for _, ingridient := range o.Ingridients {
		recipeIngridient := RecipeIngridient{}
		recipeIngridient.IDRecipe = o.ID
		recipeIngridient.IDIngridient = ingridient.ID
		recipeIngridient.Quantity = ingridient.Quantity
		recipeIngridient.UpdateById(ctx)
	}
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *Recipe) DeleteById(ctx helpers.Context) map[string]interface{} {
	id := helpers.Convert(o.ID).String()
	getById := o.GetById(ctx, id, map[string][]string{})
	ids := []int{}
	if ri, ok := getById["recipe_ingridients"].([]map[string]interface{}); ok {
		if len(ri) > 0 {
			for _, i := range ri {
				ids = append(ids, helpers.Convert(i["id"].(float64)).Int())
			}
		}
	}
	helpers.GetDB(ctx).Model(Recipe{}).Where("id IN (?)", ids).Delete(&RecipeIngridient{})
	helpers.GetDB(ctx).Model(Recipe{}).Where("id = ?", o.ID).Delete(&Recipe{})
	return helpers.DeletedMessage("recipes", "id", id)
}

func (o *Recipe) SetDefaultValue(ctx helpers.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	return params, nil
}
