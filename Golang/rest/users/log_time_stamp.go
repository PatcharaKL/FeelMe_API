package users

import (
	"net/http"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type LogTimeStamp struct {
	Id               int     `json:"id" query:"id"`
	UserName         string  `json:"username" query:"username"`
	LogTimeStampType int     `json:"log_timestamp_type" query:"log_timestamp_type"`
	UserId           int     `json:"user_id" query:"user_id"`
	Time             []uint8 `json:"time" query:"time"`
}

func (h *Handler) CheckIn(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	ac := new(models.Account)
	lgTime := new(LogTimeStamp)
	row := h.DB.QueryRow(getUserFullNameByUserId, claims.AccountId)
	if err := row.Scan(&ac.Name, &ac.Surname); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	times := time.Now().In(location)
	fullName := ac.Name + " " + ac.Surname
	if err := h.DB.QueryRow(createdLogTimeStamp, fullName, 1, claims.AccountId, times).Scan(&lgTime.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": true, "message": "Success", "Time": times})
}
func (h *Handler) CheckOut(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	ac := new(models.Account)
	lgTime := new(LogTimeStamp)
	row := h.DB.QueryRow(getUserFullNameByUserId, claims.AccountId)
	if err := row.Scan(&ac.Name, &ac.Surname); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	times := time.Now().In(location)
	fullName := ac.Name + " " + ac.Surname
	if err := h.DB.QueryRow(createdLogTimeStamp, fullName, 2, claims.AccountId, times).Scan(&lgTime.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": true, "message": "Success", "Time": times})
}
