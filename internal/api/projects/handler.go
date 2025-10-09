package projects

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ById() echo.HandlerFunc
	List() echo.HandlerFunc
}
