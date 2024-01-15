package main

import (
	"fmt"
	"log"
	"ms-paylater/config"
	"ms-paylater/docs"
	"ms-paylater/entity"
	"ms-paylater/handler"
	"ms-paylater/repository"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @contact.name Nafisa Alfiani
// @contact.email nafisa.alfiani.ica@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// init config
	cfg, err := config.InitEnv()
	if err != nil {
		log.Fatalln(err)
	}

	// init logger
	logger, err := config.InitLogger(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// init DB connection
	db, err := config.InitSql(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// auto migrate DB changes
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Transaction{},
		&entity.Loan{},
	); err != nil {
		log.Fatalln(err)
	}

	// init validator
	validator := validator.New(validator.WithRequiredStructEnabled())

	// init repository
	repo := repository.InitRepository(db)

	// init handler
	handler := handler.Init(cfg, repo, validator, logger)

	// init echo instance
	e := echo.New()

	// e.Use(middleware.Recover())
	e.Use(handler.MiddlewareLogging)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	docs.SwaggerInfo.Title = "MS Paylater"
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.GET("/ping", handler.Ping)

	v1 := e.Group("/v1/ms-paylater")
	v1.POST("/register", handler.Register)
	v1.POST("/login", handler.Login)
	v1.GET("", handler.GetUser, handler.Authorize)

	v1.POST("/loan", handler.CreateLoan, handler.Authorize)
	v1.GET("/limit", handler.GetLimit, handler.Authorize)
	v1.POST("/tarik-saldo", handler.Withdraw, handler.Authorize)
	v1.POST("/pay", handler.PayLoan, handler.Authorize)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", cfg.Server.Base, cfg.Server.Port)))
}
