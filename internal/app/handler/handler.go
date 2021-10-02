package handler

import (
	"crudTaskEcho/internal/app/model"
	"crudTaskEcho/internal/app/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//api handler structure, it handles http statuses
type CRUDHandler struct {
	service service.CRUDService
}

//constructor creates new instance of CRUDHandler using CRUDService object
func NewCRUDHandler(service service.CRUDService) *CRUDHandler {
	return &CRUDHandler{service: service}
}

func (C CRUDHandler) CreateUser(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res, err := C.service.CreateUser(c.Request().Context(), &req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (C CRUDHandler) GetUserByUUID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, userErr := C.service.GetUserById(c.Request().Context(), id)
	if userErr != nil {
		return userErr
	}
	return c.JSON(http.StatusOK, &user)
}

func (C CRUDHandler) UpdateUserByUUID(c echo.Context) error {
	var req model.UpdateUserReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return C.service.UpdateUser(c.Request().Context(), &req)
}
