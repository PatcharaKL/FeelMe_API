package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}
