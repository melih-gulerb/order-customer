package handlers

import (
	"customer/internal/handlers/types"
	"customer/internal/storages"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	r storages.Repository
}

func NewHandler(echo *echo.Group, r storages.Repository) *Handler {
	h := Handler{r: r}

	echo.POST("/create", h.Create)

	return &h

}

// Create godoc
// @Summary      Creates a customer
// @Tags         customer
// @Accept       json
// @Param 		 request body types.CreateRequest true " "
// @Success      201
// @Router       /create [post]
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

	err = h.r.InsertOne(storages.BuildCreateCustomer(request))
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, nil)
}
