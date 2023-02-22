package users

import (
	"net/http"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/labstack/echo/v4"
)

type User struct {
	AccountId    int    `json:"account_id" query:"account_id"`
	Name         string `json:"name" query:"name"`
	Surname      string `json:"surname" query:"surname"`
	Hp           int    `json:"hp" query:"hp"`
	Level        int    `json:"level" query:"level"`
	AvatarUrl    string `json:"avatar_url" query:"avatar_url"`
	PositionName string `json:"position_name" query:"position_name"`
}

func (h *Handler) GetAllUserHandler(c echo.Context) error {
	ac := new(models.Account)
	ps := new(models.Position)
	var data []User
	user := new(User)
	rows, err := h.DB.Query(getUserDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	for rows.Next() {
		if err := rows.Scan(&ac.AccountId, &ac.Name, &ac.Surname, &ac.Hp, &ac.Level, &ac.AvatarUrl, &ps.PositionName); err != nil {
			return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
		}
		user.AccountId = ac.AccountId
		user.AvatarUrl = ac.AvatarUrl.String
		user.Name = ac.Name
		user.Surname = ac.Surname
		user.Hp = ac.Hp
		user.Level = ac.Level
		user.PositionName = ps.PositionName
		data = append(data, *user)
	}
	return c.JSON(http.StatusOK, data)
}
