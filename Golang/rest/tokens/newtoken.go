package tokens

import (
	"net/http"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NewTokenHandler(c echo.Context) error {
	ac := new(models.Account)
	reToken := new(Refreshtoken)
	if err := c.Bind(reToken); err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	tk := new(Token)
	row := h.DB.QueryRow(getRefresh, reToken.Refreshtoken, true, time.Now())
	if err := row.Scan(&tk.Refreshtokens, &tk.AccountId, &tk.Exp, &tk.IsValid); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	row = h.DB.QueryRow(getAccountById, tk.AccountId)
	if err := row.Scan(&ac.AccountId, &ac.Email, &ac.PasswordHash, &ac.Name,
		&ac.Surname, &ac.AvatarUrl, &ac.ApplyDate, &ac.IsActive, &ac.Hp,
		&ac.Level, &ac.Created, &ac.DepartmentId, &ac.PositionId, &ac.CompanyId); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	refreshToken := GeneraterRefreshToken()
	token, _ := GeneraterTokenAccess(*ac)
	stmt, err := h.DB.Prepare(updateStatusRefreshToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if _, err := stmt.Exec(false, tk.Refreshtokens); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	ckRefreshToken := ""
	if err := h.DB.QueryRow(createRefreshToken, refreshToken, ac.AccountId, time.Now().Add(time.Hour*360), true).Scan(&ckRefreshToken); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"accessToken":  token,
		"refreshToken": refreshToken,
	})
}
