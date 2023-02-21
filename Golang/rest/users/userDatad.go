package users

import (
	"net/http"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/labstack/echo/v4"
)

type UserDetailForHR struct {
	Users []User `json:"userList"`
}

type User struct {
	AccountId int    `json:"account_id" query:"account_id"`
	Email     string `json:"email" query:"email"`
	Name      string `json:"name" query:"name"`
	Surname   string `json:"surname" query:"surname"`
	Hp        int    `json:"hp" query:"hp"`
	Level     int    `json:"level" query:"level"`
}

func (h *Handler) GetAllUserHandler(c echo.Context) error {
	ac := new(models.Account)
	var data UserDetailForHR
	user := new(User)
	rows, err := h.DB.Query(getUserDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	for rows.Next() {
		if err := rows.Scan(&ac.AccountId, &ac.Email, &ac.PasswordHash, &ac.Name,
			&ac.Surname, &ac.AvatarUrl, &ac.ApplyDate, &ac.IsActive, &ac.Hp,
			&ac.Level, &ac.Created, &ac.DepartmentId, &ac.PositionId, &ac.CompanyId); err != nil {
			return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
		}
		user.AccountId = ac.AccountId
		user.Email = ac.Email
		user.Name = ac.Name
		user.Surname = ac.Surname
		user.Hp = ac.Hp
		user.Level = ac.Level
		data.Users = append(data.Users, *user)
	}

	return c.JSON(http.StatusOK, data)
}
