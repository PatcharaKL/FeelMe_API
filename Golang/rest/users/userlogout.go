package users

import (
	"net/http"

	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UserLogOutHandler(c echo.Context) error {
	token := new(tokens.Refreshtoken)
	if err := c.Bind(&token); err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	stml, err := h.DB.Prepare(updateStatusRefreshToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if _, err := stml.Exec(false, token.Refreshtoken); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "")
}
