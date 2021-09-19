package service

import (
	"crudTaskEcho/internal/app/model"
	"github.com/labstack/echo/v4"
)

type CRUDService interface {
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	GetUserById(c echo.Context) (*model.GetUserRes, error)
}
