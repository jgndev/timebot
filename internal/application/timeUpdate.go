package application

import (
	"github.com/jgndev/timebot/internal/views/cards"
	"github.com/labstack/echo/v4"
)

func TimeUpdate(c echo.Context) error {
	return cards.TimeUpdate().Render(c.Request().Context(), c.Response().Writer)
}
