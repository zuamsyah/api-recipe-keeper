package controllers

import (
	"github.com/labstack/echo/v4"

	"api-recipe-keeper/helpers"
	"api-recipe-keeper/models"
)

// @Summary Get food category by id
// @Description Get food category by id
// @Tags Province
// @Accept json
// @Produce json
// @Param id path string true "Province ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Ingridient
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Failure 404 {object} helpers.HTTPNotFound
// @Router /api/food_categories/{id} [get]
func IngridientGetByIdApiHandle(c echo.Context) error {
	m := models.Ingridient{}
	res := m.GetById(helpers.SetContext(c), c.Param("id"), c.QueryParams())
	return helpers.Response(c, 200, res)
}

// @Summary Get food category list
// @Description Get food category list
// @Tags Province
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param search query string false "Search conditions"
// @Param sorts query string false "Sort by fields"
// @Param page query int false "Specify the page of results to return"
// @Param per_page query int false "Specify the number of records to return in one request"
// @Success 200 {object} helpers.HTTPList{results=[]models.Ingridient}
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Router /api/food_categories [get]
func IngridientGetPaginatedApiHandle(c echo.Context) error {
	m := models.Ingridient{}
	res := m.GetPaginated(helpers.SetContext(c), c.QueryParams())
	return helpers.Response(c, 200, res)
}

// @Summary Create new food category
// @Description Create new food category
// @Tags Province
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.Ingridient true "body"
// @Success 201 {object} models.Ingridient
// @Failure 400 {object} helpers.HTTPBadRequest
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Router /api/food_categories [post]
func IngridientCreateApiHandle(c echo.Context) error {
	o := new(models.Ingridient)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	res := o.Create(helpers.SetContext(c))
	return helpers.Response(c, 201, res)
}

// @Summary Update food category by id
// @Description Update food category by id
// @Tags Province
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.Ingridient true "body"
// @Success 200 {object} models.Ingridient
// @Failure 400 {object} helpers.HTTPBadRequest
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Router /api/food_categories/{id} [put]
func IngridientUpdateByIdApiHandle(c echo.Context) error {
	o := new(models.Ingridient)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	o.ID = helpers.Convert(c.Param("id")).Int()
	res := o.UpdateById(helpers.SetContext(c))
	return helpers.Response(c, 200, res)
}

// @Summary Partial update food category by id
// @Description Partial update food category by id
// @Tags Province
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.Ingridient true "body"
// @Success 200 {object} models.Ingridient
// @Failure 400 {object} helpers.HTTPBadRequest
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Router /api/food_categories/{id} [patch]
// func IngridientPartialUpdateByIdApiHandle(c echo.Context) error {
// 	return ProvinceUpdateByIdApiHandle(c)
// }

// @Summary Delete food category by id
// @Description Delete food category by id
// @Tags Province
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.Ingridient true "body"
// @Success 200 {object} helpers.HTTPDeleted
// @Failure 400 {object} helpers.HTTPBadRequest
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Router /api/food_categories/{id} [delete]
func IngridientDeleteByIdApiHandle(c echo.Context) error {
	o := new(models.Ingridient)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	o.ID = helpers.Convert(c.Param("id")).Int()
	res := o.DeleteById(helpers.SetContext(c))
	return helpers.Response(c, 200, res)
}
