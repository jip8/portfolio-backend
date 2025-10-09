package about

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Update() 	echo.HandlerFunc
	Get() 		echo.HandlerFunc
}
