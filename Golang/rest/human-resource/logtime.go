package humanresource

import (
	"net/http"

	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Record struct {
	ClockIn  string `json:"clock-in"`
	ClockOut string `json:"clock-out"`
}
type Result struct {
	Name    string   `json:"name"`
	Records []Record `json:"record"`
}

func (h *Handler) GetCheckInAndOut(c echo.Context) error {
	var list_cl_in []string
	var list_cl_out []string
	accountId := c.QueryParam("account-id")
	record := new(Record)
	result := new(Result)
	count := 0
	userName := ""
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	if accountId == "" {
		return c.String(http.StatusBadRequest, "")
	}
	if claims.Role != 4 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Status":  false,
			"Message": "You Not HR",
		})
	}
	rows, err := h.DB.Query(getLogTimeStamp, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	for rows.Next() {
		data := new(TimeStamp)
		if err := rows.Scan(&data.NameUser, &data.LogType, &data.Time); err != nil {
			c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		if userName == "" {
			userName = data.NameUser
		}
		if data.LogType == "Check In" {
			list_cl_in = append(list_cl_in, string(data.Time))
		}
		if data.LogType == "Check Out" {
			list_cl_out = append(list_cl_out, string(data.Time))
		}
	}
	result.Name = userName
	if len(list_cl_in) != len(list_cl_out) {
		count = len(list_cl_out)
		for i := 0; i < len(list_cl_in); i++ {
			if i < count {
				record.ClockIn = list_cl_in[i]
				record.ClockOut = list_cl_out[i]
				result.Records = append(result.Records, *record)
			}
			if i >= count {
				record.ClockIn = list_cl_in[i]
				record.ClockOut = "-"
				result.Records = append(result.Records, *record)
			}
		}
	}
	if len(list_cl_in) == len(list_cl_out) {
		for i := 0; i < len(list_cl_in); i++ {
			record.ClockIn = list_cl_in[i]
			record.ClockOut = list_cl_out[i]
			result.Records = append(result.Records, *record)
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Data": result,
	})
}
