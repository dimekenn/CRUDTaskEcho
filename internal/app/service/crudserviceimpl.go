package service

import (
	"crudTaskEcho/internal/app/model"
	"crudTaskEcho/internal/app/repository"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CRUDServiceImpl struct {
	repo repository.CRUDRepository
}

func NewCRUDService(repo repository.CRUDRepository) CRUDService {
	return &CRUDServiceImpl{repo: repo}
}

func (C CRUDServiceImpl) CreateUser(c echo.Context) error {
	var req model.GetUserRes
	if err := c.Bind(&req); err!=nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return C.repo.CreateUser(c.Request().Context(), &req)
}

func (C CRUDServiceImpl) UpdateUser(c echo.Context) error {
	var req model.UpdateUserReq
	if err := c.Bind(&req); err!=nil{
		return fmt.Errorf("bad request: %v", err.Error())
	}
	return C.repo.UpdateUser(c.Request().Context(), &req)
}

func (C CRUDServiceImpl) GetUserById(c echo.Context) (*model.GetUserRes, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		return nil, fmt.Errorf("bad request: %v", err.Error())
	}
	return C.repo.GetUserById(c.Request().Context(), id)
}

