package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateLoan create loan data for logged in user
//
// @Summary Create loan data
// @Description Create loan data for logged in user
// @Tags loan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/loan [post]
func (h *Handler) CreateLoan(c echo.Context) error {

	return h.httpSuccess(c, http.StatusOK, nil)
}

// CreateLoan create loan data for logged in user
//
// @Summary Create loan data
// @Description Create loan data for logged in user
// @Tags loan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/limit [get]
func (h *Handler) GetLimit(c echo.Context) error {

	return h.httpSuccess(c, http.StatusOK, nil)
}

// CreateLoan create loan data for logged in user
//
// @Summary Create loan data
// @Description Create loan data for logged in user
// @Tags loan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/tarik-saldo [post]
func (h *Handler) Withdraw(c echo.Context) error {

	return h.httpSuccess(c, http.StatusOK, nil)
}

// CreateLoan create loan data for logged in user
//
// @Summary Create loan data
// @Description Create loan data for logged in user
// @Tags loan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/pay [post]
func (h *Handler) PayLoan(c echo.Context) error {

	return h.httpSuccess(c, http.StatusOK, nil)
}
