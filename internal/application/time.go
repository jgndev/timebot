package application

import (
	"github.com/jgndev/timebot/internal/views/pages"
	"github.com/labstack/echo/v4"
)

func TimeUpdate(c echo.Context) error {
	return pages.TimeUpdates().Render(c.Request().Context(), c.Response().Writer)
}
