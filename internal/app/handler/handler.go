package handler

import (
	"crudTaskEcho/internal/app/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

//api handler structure, it handles http statuses
type CRUDHandler struct{
	service service.CRUDService
}

//constructor creates new instance of CRUDHandler using CRUDService object
func NewCRUDHandler(service service.CRUDService) *CRUDHandler {
	return &CRUDHandler{service: service}
}

func (C CRUDHandler) CreateUser(c echo.Context) error {
	fmt.Println("create user")
	return C.service.CreateUser(c)
}

func (C CRUDHandler) GetUserByUUID(c echo.Context) error {
	user, userErr := C.service.GetUserById(c)
	if userErr != nil{
		return userErr
	}
	return c.JSON(http.StatusOK, &user)
}

func (C CRUDHandler) UpdateUserByUUID(c echo.Context) error {
	return C.service.UpdateUser(c)
}

//func handlerHTTPStatus(err error) error  {
//	if err != nil{
//		if strings.Contains(err.Error(), "user is not found"){
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		} else if strings.Contains(err.Error(), "bad request"){
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		} else {
//			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
//		}
//	} else {
//		return nil
//	}
//
//}

