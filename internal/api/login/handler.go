package login

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Login() echo.HandlerFunc
}