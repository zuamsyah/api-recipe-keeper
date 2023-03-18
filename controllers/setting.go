package controllers

import (
	"github.com/labstack/echo/v4"

	"api-recipe-keeper/helpers"
	"api-recipe-keeper/models"
)

// @Summary Get current setting
// @Description Get current setting
// @Tags Setting
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.SettingShort
// @Failure 400 {object} helpers.HTTPBadRequest
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Router /api/setting [get]
func SettingGetApiHandle(c echo.Context) error {
	m := models.Setting{}
	return c.JSON(200, m.Get(helpers.SetContext(c)))
}

// @Summary Update current setting
// @Description Update current setting
// @Tags Setting
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.SettingShort true "body"
// @Success 200 {object} models.SettingShort
// @Failure 400 {object} helpers.HTTPBadRequest
// @Failure 401 {object} helpers.HTTPUnauthorized
// @Failure 403 {object} helpers.HTTPForbidden
// @Router /api/setting [post]
func SettingUpdateApiHandle(c echo.Context) error {
	o := echo.Map{}
	if err := c.Bind(&o); err != nil {
		return err
	}
	m := models.Setting{}
	return c.JSON(200, m.Update(helpers.SetContext(c), o))
}
