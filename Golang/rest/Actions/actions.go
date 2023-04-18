package action

import (
	"net/http"
	"strconv"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/PatcharaKL/FeelMe_API/rest/users"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Values  string `json:"values"`
}

func (h *Handler) AttackDamage(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	userId := claims.AccountId
	user_detail := new(users.User)
	ac := new(models.Account)
	log := new(AttackDamageSender)
	log_id := 0
	if err := c.Bind(log); err != nil {
		res := Response{
			Message: "Fail",
			Status:  false,
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	if log.Amount <= 0 {
		res := Response{
			Message: "Fail",
			Status:  false,
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	row := h.DB.QueryRow(users.GetUserByUserId, userId)
	if err := row.Scan(&user_detail.AccountId, &user_detail.Name, &user_detail.Surname,
		&user_detail.Hp, &user_detail.Level, &ac.AvatarUrl, &user_detail.PositionName,
		&user_detail.DepartmentName, &user_detail.CompanyName); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	if user_detail.Hp == 0 {
		res := Response{
			Message: "Fail",
			Status:  false,
			Values:  "Residual Hp of " + user_detail.Name + " " + user_detail.Surname + " = 0",
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	if user_detail.Hp-log.Amount <= 0 {
		stmt, err := h.DB.Prepare(users.UpdateUserData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		if _, err := stmt.Exec(0, userId); err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		if err := h.DB.QueryRow(createLogAttackDamage, log.Amount, time.Now(), userId).Scan(&log_id); err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		res := Response{
			Message: "Success",
			Status:  true,
			Values:  "Residual Hp of " + user_detail.Name + " " + user_detail.Surname + " = 0",
		}
		return c.JSON(http.StatusOK, res)
	}
	if user_detail.Hp-log.Amount >= 0 {
		stmt, err := h.DB.Prepare(users.UpdateUserData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		if _, err := stmt.Exec(user_detail.Hp-log.Amount, userId); err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
	}
	if err := h.DB.QueryRow(createLogAttackDamage, log.Amount, time.Now(), userId).Scan(&log_id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	res := Response{
		Message: "Success",
		Status:  true,
		Values:  "Residual Hp of " + user_detail.Name + " " + user_detail.Surname + " = " + strconv.Itoa(user_detail.Hp-log.Amount),
	}
	return c.JSON(http.StatusOK, res)
}
