package handler

import (
	"ms-paylater/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUsers returns logged in user detail
//
// @Summary Fetch user detail
// @Description Get logged in user detail
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater [get]
func (h *Handler) GetUser(c echo.Context) error {
	loggedInUsername := c.Request().Context().Value(contextKeyUsername).(string)
	user, err := h.user.Get(loggedInUsername)
	if err != nil {
		return h.httpError(c, err)
	}
	h.logger.Debug(entity.User{Username: loggedInUsername})

	return h.httpSuccess(c, http.StatusOK, user)
}
