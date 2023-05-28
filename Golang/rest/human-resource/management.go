package humanresource

import (
	"net/http"

	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type TimeStamp struct {
	NameUser string  `json:"nameUser"`
	LogType  string  `json:"logType"`
	Time     []uint8 `json:"timestamp"`
}
type TimeStampData struct {
	NameUser string `json:"nameUser"`
	LogType  string `json:"logType"`
	Time     string `json:"timestamp"`
}

const YYYYMMDD = "2006-01-02"

func (h *Handler) EditProfileEmployee(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	if claims.Role != 4 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Status":  false,
			"Message": "You Not HR",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Status":  true,
		"Message": "Hello HR",
	})
}

func (h *Handler) CreatedUser(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	if claims.Role != 4 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Status":  false,
			"Message": "You Not HR",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Status":  true,
		"Message": "Hello HR",
	})
}
