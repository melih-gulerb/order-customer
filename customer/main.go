package main

import (
	"context"
	"customer/configs"
	_ "customer/docs"
	"customer/internal/handlers"
	"customer/internal/storages"
	"customer/middlewares"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := configs.SetConfig()

	m, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoConnection))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := m.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repository := storages.NewRepository(m.Database(cfg.DBName).Collection(cfg.CustomerCollection))

	e := echo.New()
	e.Use(middlewares.PanicExceptionHandling())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "Ok!")
	})

	handlers.NewHandler(e.Group("/customer"), repository)

	e.Logger.Fatal(e.Start(":4000"))
}
