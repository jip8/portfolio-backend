package skills

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	List() 		echo.HandlerFunc
	Update() 	echo.HandlerFunc
}
