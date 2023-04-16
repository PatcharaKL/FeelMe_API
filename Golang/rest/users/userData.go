package users

import (
	"net/http"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/labstack/echo/v4"
)

type User struct {
	AccountId      int    `json:"account_id" query:"account_id"`
	Name           string `json:"name" query:"name"`
	Surname        string `json:"surname" query:"surname"`
	Hp             int    `json:"hp" query:"hp"`
	Level          int    `json:"level" query:"level"`
	AvatarUrl      string `json:"avatar_url" query:"avatar_url"`
	PositionName   string `json:"position_name" query:"position_name"`
	DepartmentName string `json:"department_name" query:"department_name"`
	CompanyName    string `json:"company_name" query:"company_name"`
}

func (h *Handler) GetAllUserHandler(c echo.Context) error {
	userId := c.QueryParam("accountId")
	ac := new(models.Account)
	ps := new(models.Position)
	dm := new(models.Department)
	cp := new(models.Company)
	var data []User
	user := new(User)
	if userId != "" {
		data := new(User)
		row := h.DB.QueryRow(getUserByUserId, userId)
		if err := row.Scan(&data.AccountId, &data.Name, &data.Surname, &data.Hp, &data.Level, &ac.AvatarUrl, &data.PositionName, &data.DepartmentName, &data.CompanyName); err != nil {
			return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
		}
		data.AvatarUrl = ac.AvatarUrl.String
		return c.JSON(http.StatusOK, data)
	}
	rows, err := h.DB.Query(getUserDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	for rows.Next() {
		if err := rows.Scan(&ac.AccountId, &ac.Name, &ac.Surname, &ac.Hp, &ac.Level, &ac.AvatarUrl, &ps.PositionName, &dm.DepartmentName, &cp.CompanyName); err != nil {
			return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
		}
		user.AccountId = ac.AccountId
		user.AvatarUrl = ac.AvatarUrl.String
		user.Name = ac.Name
		user.Surname = ac.Surname
		user.Hp = ac.Hp
		user.Level = ac.Level
		user.PositionName = ps.PositionName
		user.DepartmentName = dm.DepartmentName
		user.CompanyName = cp.CompanyName
		data = append(data, *user)
	}
	return c.JSON(http.StatusOK, data)
}
