package handler

import (
	"ms-paylater/entity"
	"ms-paylater/errors"
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
// @Param loan_request body entity.LoanRequest true "loan request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/loan [post]
func (h *Handler) CreateLoan(c echo.Context) error {
	loanReq := entity.LoanRequest{}
	if err := c.Bind(&loanReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(loanReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	tier, limit := loanReq.AssignTierAndLimit()
	loan := entity.Loan{
		UserId:  int(c.Request().Context().Value(contextKeyUserId).(float64)),
		Tier:    tier,
		Limit:   limit,
		Balance: limit,
	}
	newLoan, err := h.loan.Create(loan)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, newLoan)
}

// GetLimit fetch user balance and limit
//
// @Summary Get user limit
// @Description Fetch user balance and limit
// @Tags loan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/limit [get]
func (h *Handler) GetLimit(c echo.Context) error {
	userId := int(c.Request().Context().Value(contextKeyUserId).(float64))
	loan, err := h.loan.Get(userId)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, loan)
}

// Withdraw log transaction to withdraw from balance
//
// @Summary Withdraw user balance
// @Description Log transaction to withdraw from balance
// @Tags loan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param trx_request body entity.TransactionRequest true "loan request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/tarik-saldo [post]
func (h *Handler) Withdraw(c echo.Context) error {
	trxReq := entity.TransactionRequest{}
	if err := c.Bind(&trxReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(trxReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	userId := int(c.Request().Context().Value(contextKeyUserId).(float64))
	transaction := entity.Transaction{
		UserId: userId,
		Amount: trxReq.Amount,
		Type:   entity.TransactionTypeDebit,
	}
	newTransaction, err := h.transaction.Create(transaction)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusCreated, newTransaction)
}

// PayLoan log transaction to pay user balance
//
// @Summary Recover user balance
// @Description Log transaction to pay user balance
// @Tags loan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param trx_request body entity.TransactionRequest true "loan request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /v1/ms-paylater/pay [post]
func (h *Handler) PayLoan(c echo.Context) error {
	trxReq := entity.TransactionRequest{}
	if err := c.Bind(&trxReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(trxReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	userId := int(c.Request().Context().Value(contextKeyUserId).(float64))
	transaction := entity.Transaction{
		UserId: userId,
		Amount: trxReq.Amount,
		Type:   entity.TransactionTypeCredit,
	}
	newTransaction, err := h.transaction.Create(transaction)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusCreated, newTransaction)
}
