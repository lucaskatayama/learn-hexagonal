package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/lucaskatayama/learn-hexagonal/go/app/service"
	"log"
	"net/http"
	"strconv"
)

type Rest struct {
	svc *service.Service
}

func New(svc *service.Service) *Rest {
	return &Rest{svc: svc}
}

func (rest *Rest) GetSomething(c echo.Context) error {
	name := c.QueryParam("name")
	ageStr := c.QueryParam("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return err
	}
	params := service.Params{
		Name: name,
		Age:  age,
	}
	res, err := rest.svc.ServiceMethod(c.Request().Context(), params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (rest *Rest) Start() {
	e := echo.New()

	e.GET("/", rest.GetSomething)

	log.Fatal(e.Start("localhost:8080"))
}
