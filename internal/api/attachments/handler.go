package attachments

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Get() echo.HandlerFunc
}
