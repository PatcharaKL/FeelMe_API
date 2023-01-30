package users

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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
func GeneraterRefreshToken() string {
	refreshToken := []string{}
	for i := 0; i < 4; i++ {
		id := uuid.New()
		refreshToken = append(refreshToken, id.String())
	}
	return strings.Join(refreshToken, "-")
}
func GeneraterTokenAccess(ac Account) (string, error) {
	claims := &JwtCustomClaims{ac.Email, ac.Name, ac.Surname, ac.PositionId, ac.AccountId, ac.DepartmentId,
		ac.CompanyId, jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("GVebOWpKrqyZ9RwPXzazpNpcmA6njskh"))
	if err != nil {
		return "", err
	}
	return t, nil
}
func (h *Handler) UserLoginHandler(c echo.Context) error {
	ac := new(Account)
	u := new(userLogin)
	err := c.Bind(u)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	user := userLoginDTO{
		Email:    u.Email,
		Password: u.Password,
	}
	row := h.DB.QueryRow(getAccountById, user.Email)
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
	token, _ := GeneraterTokenAccess(*ac)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	refreshToken := GeneraterRefreshToken()
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
		err := rows.Scan(&tk)
		if err != nil {
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
