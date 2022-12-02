package main

import (
	"context"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"order/configs"
	_ "order/docs"
	"order/internal/clients"
	"order/internal/handlers"
	"order/internal/storages"
	"order/middlewares"
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

	e := echo.New()
	e.Use(middlewares.PanicExceptionHandling())

	repository := storages.NewRepository(m.Database(cfg.DBName).Collection(cfg.OrderCollection))
	customerClient := clients.NewCustomerClient(cfg.CustomerClient)

	handlers.NewHandler(e.Group("/order"), repository, *customerClient)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "Ok!")
	})

	e.Logger.Fatal(e.Start(":2000"))
}
