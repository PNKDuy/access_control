package controller

import (
	"access_control/model/apidetail"
	"access_control/model/casbin"
	"access_control/model/request"
	"access_control/model/role"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CheckPermission check perrmission
// @Summary manipulate message from Kafka
// @tags access-control
// @Param model_type body object true "model_type"
// @Produce json
// @Success 200 {object} request.Request
// @Failure 400 {HTTPError} HTTPError
// @Router /access-control [post]
func CheckPermission(c echo.Context) (err error) {
	req := request.Request{}
	if err = c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	msg, err := req.CheckPermissionService()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSONPretty(http.StatusOK, &msg, " ")
}

// Create create object
// @Summary Create Object with Specified Type
// @Tags general
// @Param type path string true "type"
// @Param model-value body object true "model-value"
// @Produce json
// @Success 200
// @Failure 400 {HTTPError} HTTPError
// @Router /general/{type} [post]
func Create(c echo.Context) (err error) {
	modelType := c.Param("type")
	switch modelType {
	case "casbin":
		{
			casbin := casbin.Casbin{}
			if err = c.Bind(&casbin); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}	
			_, err = casbin.Create()
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.NoContent(http.StatusNoContent)
		}
	case "role":
		{
			role := role.Role{}
			if err = c.Bind(&role); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			_, err = role.Create()
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.NoContent(http.StatusNoContent)
		}
	case "api_detail":
		{
			apidetail := apidetail.ApiDetail{}
			if err = c.Bind(&apidetail); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			_, err = apidetail.Create()
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.NoContent(http.StatusNoContent)
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("model type does not exist"))
	}
}

// Get Get objects by limit
// @Summary Show the list of objects by limit input.
// @Tags general
// @Param type path string true "type"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {HTTPError} HTTPError
// @Router /general/{type} [get]
func Get(c echo.Context) (err error) {
	modelType := c.Param("type")

	switch modelType {
	case "role":
		{
			roles, err := role.Get()
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSONPretty(http.StatusOK, roles, " ")
		}
	case "api_detail":
		{
			apiDetails, err := apidetail.Get()
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSONPretty(http.StatusOK, apiDetails, " ")
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("model type does not exist"))

	}
}

// GetById Get Object by Id
// @Summary Get active object by Id
// @Tags general
// @Param type path string true "type"
// @Param id path string true "id"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {HTTPError} HTTPError
// @Router /general/{type}/{id} [get]
func GetById(c echo.Context) (err error) {
	modelType := c.Param("type")
	id := c.Param("id")

	switch modelType {
	case "role":
		{
			role, err := role.GetById(id)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			return c.JSONPretty(http.StatusOK, &role, "  ")
		}
	case "api_detail":
		{
			apidetail, err := apidetail.GetById(id)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSONPretty(http.StatusOK, &apidetail, "  ")
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("model type does not exist"))
	}
}

// Update Update specified object with id
// @Summary Update specified object with id
// @Tags general
// @Param type path string true "type"
// @Param id path string true "id"
// @Param model_value body object true "model_value"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {HTTPError} HTTPError
// @Router /general/{type}/{id} [put]
func Update(c echo.Context) (err error) {
	modelType := c.Param("type")
	id := c.Param("id")

	switch modelType {
	case "role":
		{
			role, err := role.GetById(id)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			if err = c.Bind(&role); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			role, err = role.Update()
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSONPretty(http.StatusOK, &role, "  ")
		}
	case "api_detail":
		{
			apiDetail, err := apidetail.GetById(id)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			if err = c.Bind(&apiDetail); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			apiDetail, err = apiDetail.Update()
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSONPretty(http.StatusOK, &apiDetail, "  ")
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("model type does not exist"))
	}
}

// Delete update is_deleted field
// @Summary Deactive object by user_id
// @Tags general
// @Param type path string true "type"
// @Param id path string true "id"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {HTTPError} HTTPError
// @Router /general/{type}/{id} [delete]
func Delete(c echo.Context) (err error) {
	modelType := c.Param("type")
	id := c.Param("id")

	switch modelType {
	case "role":
		{
			err := role.Delete(id)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.NoContent(http.StatusNoContent)
		}
	case "api_detail":
		{
			err := apidetail.Delete(id)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.NoContent(http.StatusNoContent)
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("model type does not exist"))
	}
}
