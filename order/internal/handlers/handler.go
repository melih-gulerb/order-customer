package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"order/internal/clients"
	"order/internal/handlers/types"
	"order/internal/storages"
)

type Handler struct {
	r              storages.Repository
	customerClient clients.CustomerClient
}

func NewHandler(echo *echo.Group, r storages.Repository, c clients.CustomerClient) *Handler {
	h := Handler{r: r, customerClient: c}

	echo.POST("/create", h.Create)

	return &h

}

func (h Handler) Create(c echo.Context) error {
	var request types.CreateRequest

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		panic(err)
	}

	if err := h.customerClient.ValidateCustomerExists(request.CustomerId); err != nil {
		panic(err)
	}

	err = h.r.InsertOne(storages.BuildCreateOrder(request))
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, nil)
}
