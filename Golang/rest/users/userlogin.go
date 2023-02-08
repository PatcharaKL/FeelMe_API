package users

import (
	"net/http"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) UserLoginHandler(c echo.Context) error {
	ac := new(models.Account)
	u := new(userLogin)
	err := c.Bind(u)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	user := userLoginDTO{
		Email:    u.Email,
		Password: u.Password,
	}
	row := h.DB.QueryRow(getAccountByEmail, user.Email)
	if err := row.Scan(&ac.AccountId, &ac.Email, &ac.PasswordHash, &ac.Name,
		&ac.Surname, &ac.AvatarUrl, &ac.ApplyDate, &ac.IsActive, &ac.Hp,
		&ac.Level, &ac.Created, &ac.DepartmentId, &ac.PositionId, &ac.CompanyId); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	err = CheckPasswordHash(user.Password, ac.PasswordHash)
	if err != nil || ac.PasswordHash == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}
	if err = UpdateStatusRefreshToken(h, ac.AccountId); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	token := tokens.GeneraterTokenAccess(*ac)
	refreshToken := tokens.GeneraterRefreshToken()
	ckRefreshToken := ""
	if err := h.DB.QueryRow(createRefreshToken, refreshToken, ac.AccountId, time.Now().Add(time.Hour*360), true).Scan(&ckRefreshToken); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"accessToken":  token,
		"refreshToken": refreshToken,
	})
}
func UpdateStatusRefreshToken(h *Handler, accountId int) error {
	rows, err := h.DB.Query(getRefreshTokenByAccountId, accountId, true)
	if err != nil {
		return err
	}
	reks := []string{}
	for rows.Next() {
		tk := ""
		if err := rows.Scan(&tk); err != nil {
			return err
		}
		reks = append(reks, tk)
	}
	for i := 0; i < len(reks); i++ {
		stmt, err := h.DB.Prepare(updateStatusRefreshToken)
		if err != nil {
			return err
		}

		if _, err := stmt.Exec(false, reks[i]); err != nil {
			return err
		}
	}
	return nil
}
