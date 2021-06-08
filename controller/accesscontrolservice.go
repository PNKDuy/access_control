package controller

import (
	"access_control/model/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CheckPermission check perrmission
// @Summary manipulate date
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
