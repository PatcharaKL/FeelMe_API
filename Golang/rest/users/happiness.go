package users

import (
	"net/http"
	"time"

	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h *Handler) HappinesspointHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	userId := claims.AccountId

	hpyPointBody := new(HapPointRequest)
	if err := c.Bind(hpyPointBody); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	hpyPoint := HappinessPoint{}
	if err := h.DB.QueryRow(createdHappinessPoint, userId, hpyPointBody.Selfpoints, hpyPointBody.Workpoints, hpyPointBody.Copoints, time.Now()).Scan(&hpyPoint.Id); err != nil {

		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"user_id":          userId,
		"self_points":      hpyPointBody.Selfpoints,
		"work_points":      hpyPointBody.Workpoints,
		"co_worker_points": hpyPointBody.Copoints,
	})
}
