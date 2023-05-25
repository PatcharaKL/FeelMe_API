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

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func (h *Handler) CheckIn(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	ac := new(models.Account)
	lgTime := new(LogTimeStamp)
	row := h.DB.QueryRow(getUserFullNameByUserId, claims.AccountId)
	if err := row.Scan(&ac.Name, &ac.Surname); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	location := time.FixedZone("UTC+7", 7*60*60)

	times := time.Now().In(location).Format(DDMMYYYYhhmmss)
	fullName := ac.Name + " " + ac.Surname
	if err := h.DB.QueryRow(createdLogTimeStamp, fullName, 1, claims.AccountId, times).Scan(&lgTime.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": true, "message": "Success", "Time": times})
}
func (h *Handler) CheckOut(c echo.Context) error {
	CheckHTTP()
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	hpyPointBody := new(HapPointRequest)
	if c.Request().Body == http.NoBody {
		return c.JSON(http.StatusBadRequest, "")
	}
	if err := c.Bind(hpyPointBody); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	ac := new(models.Account)
	lgTime := new(LogTimeStamp)
	row := h.DB.QueryRow(getUserFullNameByUserId, claims.AccountId)
	if err := row.Scan(&ac.Name, &ac.Surname); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	location := time.FixedZone("UTC+7", 7*60*60)

	times := time.Now().In(location).Format(DDMMYYYYhhmmss)
	fullName := ac.Name + " " + ac.Surname
	if err := h.DB.QueryRow(createdLogTimeStamp, fullName, 2, claims.AccountId, times).Scan(&lgTime.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	stmt, err := h.DB.Prepare(UpdateUserData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if _, err := stmt.Exec(100, claims.AccountId); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	hpyPoint := models.HappinessPoint{}
	if err := h.DB.QueryRow(createdHappinessPoint, claims.AccountId, hpyPointBody.Selfpoints, hpyPointBody.Workpoints, hpyPointBody.Copoints, times).Scan(&hpyPoint.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	value_over_all, err := FuzzyCalculatorAll(hpyPointBody.Selfpoints, hpyPointBody.Workpoints, hpyPointBody.Copoints)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	fuzzy_self_points, err := FuzzyCalculator(hpyPointBody.Selfpoints)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	fuzzy_work_points, err := FuzzyCalculator(hpyPointBody.Workpoints)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	co_worker_points, err := FuzzyCalculator(hpyPointBody.Copoints)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	id := 0
	if err := h.DB.QueryRow(createFuzzyValue, fuzzy_self_points.Value, fuzzy_work_points.Value, co_worker_points.Value, value_over_all.Value, times, claims.AccountId).Scan(&id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": true, "message": "Success", "Time": times})
}
