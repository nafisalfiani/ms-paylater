package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"ms-paylater/config"
	"ms-paylater/entity"
	"ms-paylater/errors"
	"ms-paylater/repository"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	config      *config.Value
	validator   *validator.Validate
	logger      *logrus.Logger
	user        repository.UserInterface
	loan        repository.LoanInterface
	transaction repository.TransactionInterface
}

// Init create new Handler object
func Init(config *config.Value, repo *repository.Repository, validator *validator.Validate, logger *logrus.Logger) *Handler {
	return &Handler{
		config:      config,
		validator:   validator,
		logger:      logger,
		user:        repo.User,
		loan:        repo.Loan,
		transaction: repo.Transaction,
	}
}

func (h *Handler) Ping(c echo.Context) error {
	return h.httpSuccess(c, http.StatusOK, "PONG!")
}

// httpError is helper function for error response
func (h *Handler) httpError(c echo.Context, err error, additionalMessage ...string) error {
	p, _ := os.Getwd()

	_, filename, line, _ := runtime.Caller(1)
	log.Printf("\033[31m[error]\033[0m \033[35m%s:%d\033[0m -> %v", strings.TrimPrefix(filename, p), line, err)

	resp := entity.HttpResp{
		Status:  errors.GetStatusCode(err),
		Message: http.StatusText(errors.GetStatusCode(err)),
		Error:   fmt.Sprintf("%s. %s", err.Error(), additionalMessage),
	}

	return h.ResponseLogging(c, errors.GetStatusCode(err), resp)
}

// httpSuccess is helper function for success response
func (h *Handler) httpSuccess(c echo.Context, statusCode int, data any) error {
	resp := entity.HttpResp{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Data:    data,
	}

	return h.ResponseLogging(c, statusCode, resp)
}

func (h *Handler) ResponseLogging(c echo.Context, code int, resp any) error {
	res, _ := json.Marshal(resp)
	trail := h.logger.WithFields(
		logrus.Fields{
			"at":   time.Now().Format(time.RFC3339),
			"resp": string(res),
		},
	)

	trail.Info("response")

	return c.JSON(code, resp)
}

func (h *Handler) MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		trail := h.logger.WithFields(
			logrus.Fields{
				"at":     time.Now().Format(time.RFC3339),
				"method": c.Request().Method,
				"uri":    c.Request().URL.String(),
				"ip":     c.Request().RemoteAddr,
			},
		)

		trail.Info("incoming request")

		return next(c)
	}
}
