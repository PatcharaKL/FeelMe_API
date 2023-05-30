package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const YYYYMMDD = "2006-01-02"

var HTTP = "http://127.0.0.1:8000/"

var check = false

type Hppoint struct {
	SelfPoints     int     `json:"self_points"`
	WorkPoints     int     `json:"work_points" `
	CoWorkerPoints int     `json:"co_worker_points"`
	FuzzyValue     float64 `json:"fuzzy_value" `
}
type Record struct {
	Hppoints Hppoint `json:"happiness_points"`
	Date     string  `json:"date"`
}
type ResponseGetHappines struct {
	Id      int      `json:"id"`
	Period  string   `json:"period"`
	Records []Record `json:"record"`
}
type Value struct {
	Value float64 `json:"value"`
}

type SP struct {
	SelfPoints []int `json:"self_hps"`
}
type FuzzyValues struct {
	Id             int     `json:"id" query:"id"`
	SelfPoints     int     `json:"fuzzy_self_points" query:"fuzzy_self_points"`
	WorkPoints     int     `json:"fuzzy_work_points" query:"fuzzy_work_points"`
	CoWorkerPoints int     `json:"fuzzy_co_worker_points" query:"fuzzy_co_worker_points"`
	ValueOverAll   int     `json:"value_over_all" query:"value_over_all"`
	Time           []uint8 `json:"timestamp" query:"timestamp"`
	AccountId      int     `json:"account_id" query:"account_id"`
}
type FuzzyValuesByPoistion struct {
	SelfPoints     int    `json:"fuzzy_self_points" query:"fuzzy_self_points"`
	WorkPoints     int    `json:"fuzzy_work_points" query:"fuzzy_work_points"`
	CoWorkerPoints int    `json:"fuzzy_co_worker_points" query:"fuzzy_co_worker_points"`
	ValueOverAll   int    `json:"value_over_all" query:"value_over_all"`
	AccountId      int    `json:"account_id" query:"account_id"`
	PositionId     int    `json:"position_id" query:"position_id"`
	PositionName   string `json:"position_name" query:"position_name"`
}
type FuzzyValuesByDepartMent struct {
	SelfPoints     int    `json:"fuzzy_self_points" query:"fuzzy_self_points"`
	WorkPoints     int    `json:"fuzzy_work_points" query:"fuzzy_work_points"`
	CoWorkerPoints int    `json:"fuzzy_co_worker_points" query:"fuzzy_co_worker_points"`
	ValueOverAll   int    `json:"value_over_all" query:"value_over_all"`
	AccountId      int    `json:"account_id" query:"account_id"`
	DepartMentId   int    `json:"department_id" query:"department_id"`
	DepartMentName string `json:"department_name" query:"department_name"`
}

func FuzzyCalculatorAll(self_points int, work_points int, co_points int) (*Value, error) {
	http_name := HTTP + fmt.Sprintf("v1/fuzzy?self_hp=%d&work_hp=%d&co_worker_hp=%d", self_points, work_points, co_points)
	vauel := new(Value)
	req, err := http.Get(http_name)
	if err != nil {
		return nil, err
	}
	json.NewDecoder(req.Body).Decode(vauel)
	return vauel, nil
}
func FuzzyCalculator(points int) (*Value, error) {
	http_name := HTTP + fmt.Sprintf("v1/fuzzy/cal?point=%d", points)
	vauel := new(Value)
	req, err := http.Get(http_name)
	log.Print(points)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	json.NewDecoder(req.Body).Decode(vauel)
	return vauel, nil
}
func (h *Handler) HappinesspointHandler(c echo.Context) error {
	CheckHTTP()
	location := time.FixedZone("UTC+7", 7*60*60)
	times := time.Now().In(location).Format(DDMMYYYYhhmmss)
	// userId := c.Param("id")
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	userId := claims.AccountId
	hpyPointBody := new(HapPointRequest)
	if err := c.Bind(hpyPointBody); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	id := 0
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
	hpyPoint := models.HappinessPoint{}
	if err := h.DB.QueryRow(createdHappinessPoint, userId, hpyPointBody.Selfpoints, hpyPointBody.Workpoints, hpyPointBody.Copoints, time.Now()).Scan(&hpyPoint.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if err := h.DB.QueryRow(createFuzzyValue, fuzzy_self_points.Value, fuzzy_work_points.Value, co_worker_points.Value, value_over_all.Value, times, userId).Scan(&id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"user_id":          userId,
		"self_points":      fuzzy_self_points.Value,
		"work_points":      fuzzy_work_points.Value,
		"co_worker_points": co_worker_points.Value,
		"value_over_all":   value_over_all.Value,
	})

}

func CheckHTTP() {
	if _, err := http.Get(HTTP); err != nil {
		check = true
		HTTP = "http://fuzzy-api:8000/"
	}
}
func (h *Handler) GetHappinessScoreAverage(c echo.Context) error {
	// period := c.QueryParam("period")
	accountId := c.QueryParam("account-id")
	var fuzzy_data []FuzzyValues
	count := 0
	fuzzy_self_points_average := 0
	fuzzy_work_points_average := 0
	fuzzy_co_worker_points_average := 0
	value_over_all_average := 0
	switch period := c.QueryParam("period"); period {
	case "month":
		{
			st1 := ""
			st2 := ""
			location := time.FixedZone("UTC+7", 7*60*60)
			times := time.Now().In(location).Format(DDMMYYYYhhmmss)
			arr := strings.Split(times, "-")
			timeStart := "2023-" + arr[1] + "-01 00:00:00"
			timeStop := "2023-" + arr[1] + "-31 23:59:59"
			if accountId != "" {
				rows, err := h.DB.Query(getHappinessScoreAllByAccountIdAndDate, accountId, timeStart, timeStop)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				fuzzy := new(FuzzyValues)
				for rows.Next() {
					if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					if st1 == "" {
						st1 = string(fuzzy.Time)
					}
					count++
					fuzzy_self_points_average += fuzzy.SelfPoints
					fuzzy_work_points_average += fuzzy.WorkPoints
					fuzzy_co_worker_points_average += fuzzy.CoWorkerPoints
					value_over_all_average += fuzzy.ValueOverAll
					fuzzy_data = append(fuzzy_data, *fuzzy)
					st2 = string(fuzzy.Time)
				}
				if len(fuzzy_data) == 0 {
					return c.JSON(http.StatusNoContent, "")
				}
				fuzzy_self_points_average = fuzzy_self_points_average / count
				fuzzy_work_points_average = fuzzy_work_points_average / count
				fuzzy_co_worker_points_average = fuzzy_co_worker_points_average / count
				value_over_all_average = value_over_all_average / count

				return c.JSON(http.StatusOK, echo.Map{
					"fuzzy_self_points_average":      fuzzy_self_points_average,
					"fuzzy_work_points_average":      fuzzy_work_points_average,
					"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
					"value_over_all_average":         value_over_all_average,
					"Time":                           st1 + " - " + st2,
				})
			}
			rows, err := h.DB.Query(getHappinessScoreAllByDate, timeStart, timeStop)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy := new(FuzzyValues)
			for rows.Next() {
				if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				if st1 == "" {
					st1 = string(fuzzy.Time)
				}
				count++
				fuzzy_self_points_average += fuzzy.SelfPoints
				fuzzy_work_points_average += fuzzy.WorkPoints
				fuzzy_co_worker_points_average += fuzzy.CoWorkerPoints
				value_over_all_average += fuzzy.ValueOverAll
				fuzzy_data = append(fuzzy_data, *fuzzy)
				st2 = string(fuzzy.Time)
			}
			if len(fuzzy_data) == 0 {
				return c.JSON(http.StatusNoContent, "")
			}
			fuzzy_self_points_average = fuzzy_self_points_average / count
			fuzzy_work_points_average = fuzzy_work_points_average / count
			fuzzy_co_worker_points_average = fuzzy_co_worker_points_average / count
			value_over_all_average = value_over_all_average / count

			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
				"Time":                           st1 + " - " + st2,
			})
		}
	case "year":
		{
			st1 := ""
			st2 := ""
			location := time.FixedZone("UTC+7", 7*60*60)
			times := time.Now().In(location).Format(DDMMYYYYhhmmss)
			arr := strings.Split(times, "-")
			timeStart := arr[0] + "-" + "01-01 00:00:00"
			timeStop := arr[0] + "-" + "12-31 23:59:59"
			if accountId != "" {
				rows, err := h.DB.Query(getHappinessScoreAllByAccountIdAndDate, accountId, timeStart, timeStop)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				fuzzy := new(FuzzyValues)
				for rows.Next() {
					if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					if st1 == "" {
						st1 = string(fuzzy.Time)
					}
					count++
					fuzzy_self_points_average += fuzzy.SelfPoints
					fuzzy_work_points_average += fuzzy.WorkPoints
					fuzzy_co_worker_points_average += fuzzy.CoWorkerPoints
					value_over_all_average += fuzzy.ValueOverAll
					fuzzy_data = append(fuzzy_data, *fuzzy)
					st2 = string(fuzzy.Time)
				}
				if len(fuzzy_data) == 0 {
					return c.JSON(http.StatusNoContent, "")
				}
				fuzzy_self_points_average = fuzzy_self_points_average / count
				fuzzy_work_points_average = fuzzy_work_points_average / count
				fuzzy_co_worker_points_average = fuzzy_co_worker_points_average / count
				value_over_all_average = value_over_all_average / count

				return c.JSON(http.StatusOK, echo.Map{
					"fuzzy_self_points_average":      fuzzy_self_points_average,
					"fuzzy_work_points_average":      fuzzy_work_points_average,
					"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
					"value_over_all_average":         value_over_all_average,
					"Time":                           st1 + " - " + st2,
				})
			}
			rows, err := h.DB.Query(getHappinessScoreAllByDate, timeStart, timeStop)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy := new(FuzzyValues)
			for rows.Next() {
				if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				if st1 == "" {
					st1 = string(fuzzy.Time)
				}
				count++
				fuzzy_self_points_average += fuzzy.SelfPoints
				fuzzy_work_points_average += fuzzy.WorkPoints
				fuzzy_co_worker_points_average += fuzzy.CoWorkerPoints
				value_over_all_average += fuzzy.ValueOverAll
				fuzzy_data = append(fuzzy_data, *fuzzy)
				st2 = string(fuzzy.Time)
			}
			if len(fuzzy_data) == 0 {
				return c.JSON(http.StatusNoContent, "")
			}
			fuzzy_self_points_average = fuzzy_self_points_average / count
			fuzzy_work_points_average = fuzzy_work_points_average / count
			fuzzy_co_worker_points_average = fuzzy_co_worker_points_average / count
			value_over_all_average = value_over_all_average / count

			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
				"Time":                           st1 + " - " + st2,
			})
		}
	case "":
		{
			if accountId != "" {
				rows, err := h.DB.Query(getHappinessScoreAllByAccountId, accountId)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				fuzzy := new(FuzzyValues)
				for rows.Next() {
					if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					count++
					fuzzy_self_points_average += fuzzy.SelfPoints
					fuzzy_work_points_average += fuzzy.WorkPoints
					fuzzy_co_worker_points_average += fuzzy.CoWorkerPoints
					value_over_all_average += fuzzy.ValueOverAll
					fuzzy_data = append(fuzzy_data, *fuzzy)
				}
				if len(fuzzy_data) == 0 {
					return c.JSON(http.StatusNoContent, "")
				}
				fuzzy_self_points_average = fuzzy_self_points_average / count
				fuzzy_work_points_average = fuzzy_work_points_average / count
				fuzzy_co_worker_points_average = fuzzy_co_worker_points_average / count
				value_over_all_average = value_over_all_average / count

				return c.JSON(http.StatusOK, echo.Map{
					"fuzzy_self_points_average":      fuzzy_self_points_average,
					"fuzzy_work_points_average":      fuzzy_work_points_average,
					"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
					"value_over_all_average":         value_over_all_average,
				})
			}
			rows, err := h.DB.Query(getHappinessScoreAll)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy := new(FuzzyValues)
			for rows.Next() {
				if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				count++
				fuzzy_self_points_average += fuzzy.SelfPoints
				fuzzy_work_points_average += fuzzy.WorkPoints
				fuzzy_co_worker_points_average += fuzzy.CoWorkerPoints
				value_over_all_average += fuzzy.ValueOverAll
				fuzzy_data = append(fuzzy_data, *fuzzy)
			}
			if len(fuzzy_data) == 0 {
				return c.JSON(http.StatusNoContent, "")
			}
			fuzzy_self_points_average = fuzzy_self_points_average / count
			fuzzy_work_points_average = fuzzy_work_points_average / count
			fuzzy_co_worker_points_average = fuzzy_co_worker_points_average / count
			value_over_all_average = value_over_all_average / count

			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	default:
		{
			return c.JSON(http.StatusBadRequest, "")
		}
	}

}

type FuzzyAverage struct {
	Value    float32
	DateTime string
}
type FuzzyAverageByPoistion struct {
	Value    float32
	Position string
}
type FuzzyAverageByDepartMent struct {
	Value      float32
	Department string
}

func (h *Handler) GetHappinessScoreAllTimeAverage(c echo.Context) error {
	var fuzzy_self_points_average []FuzzyAverage
	var fuzzy_work_points_average []FuzzyAverage
	var fuzzy_co_worker_points_average []FuzzyAverage
	var value_over_all_average []FuzzyAverage
	location := time.FixedZone("UTC+7", 7*60*60)
	fuzzy_self_points := 0
	fuzzy_work_points := 0
	fuzzy_co_worker_points := 0
	value_over_all := 0
	count := 0
	var srtDate [2]string
	startDate := ""
	accountId := c.QueryParam("account-id")
	switch period := c.QueryParam("period"); period {
	case "":
		{
			if accountId != "" {
				rows, err := h.DB.Query(getHappinessScoreAllByAccountId, accountId)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				fuzzy := new(FuzzyValues)
				for rows.Next() {
					if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					res := strings.Split(string(fuzzy.Time), " ")
					if startDate == "" {
						startDate = res[0]
					}
					if startDate == res[0] {
						count++
						fuzzy_self_points += fuzzy.SelfPoints
						fuzzy_work_points += fuzzy.WorkPoints
						fuzzy_co_worker_points += fuzzy.CoWorkerPoints
						value_over_all += fuzzy.ValueOverAll
					}
					if startDate != res[0] {
						data1 := new(FuzzyAverage)
						data2 := new(FuzzyAverage)
						data3 := new(FuzzyAverage)
						data4 := new(FuzzyAverage)
						data1.Value = float32(fuzzy_self_points) / float32(count)
						data2.Value = float32(fuzzy_work_points) / float32(count)
						data3.Value = float32(fuzzy_co_worker_points) / float32(count)
						data4.Value = float32(value_over_all) / float32(count)
						data1.DateTime = startDate
						data2.DateTime = startDate
						data3.DateTime = startDate
						data4.DateTime = startDate
						fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
						fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
						fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
						value_over_all_average = append(value_over_all_average, *data4)
						fuzzy_self_points = fuzzy.SelfPoints
						fuzzy_work_points = fuzzy.WorkPoints
						fuzzy_co_worker_points = fuzzy.CoWorkerPoints
						value_over_all = fuzzy.ValueOverAll
						count = 1
						startDate = res[0]
					}
				}
				data1 := new(FuzzyAverage)
				data2 := new(FuzzyAverage)
				data3 := new(FuzzyAverage)
				data4 := new(FuzzyAverage)
				data1.Value = float32(fuzzy_self_points) / float32(count)
				data2.Value = float32(fuzzy_work_points) / float32(count)
				data3.Value = float32(fuzzy_co_worker_points) / float32(count)
				data4.Value = float32(value_over_all) / float32(count)
				data1.DateTime = startDate
				data2.DateTime = startDate
				data3.DateTime = startDate
				data4.DateTime = startDate
				fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
				fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
				fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
				value_over_all_average = append(value_over_all_average, *data4)
				return c.JSON(http.StatusOK, echo.Map{
					"fuzzy_self_points_average":      fuzzy_self_points_average,
					"fuzzy_work_points_average":      fuzzy_work_points_average,
					"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
					"value_over_all_average":         value_over_all_average,
				})
			}
			rows, err := h.DB.Query(getHappinessScoreAll)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy := new(FuzzyValues)
			for rows.Next() {
				if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				res := strings.Split(string(fuzzy.Time), " ")
				if startDate == "" {
					startDate = res[0]
				}
				if startDate == res[0] {
					count++
					fuzzy_self_points += fuzzy.SelfPoints
					fuzzy_work_points += fuzzy.WorkPoints
					fuzzy_co_worker_points += fuzzy.CoWorkerPoints
					value_over_all += fuzzy.ValueOverAll
				}
				if startDate != res[0] {
					data1 := new(FuzzyAverage)
					data2 := new(FuzzyAverage)
					data3 := new(FuzzyAverage)
					data4 := new(FuzzyAverage)
					data1.Value = float32(fuzzy_self_points) / float32(count)
					data2.Value = float32(fuzzy_work_points) / float32(count)
					data3.Value = float32(fuzzy_co_worker_points) / float32(count)
					data4.Value = float32(value_over_all) / float32(count)
					data1.DateTime = startDate
					data2.DateTime = startDate
					data3.DateTime = startDate
					data4.DateTime = startDate
					fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
					fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
					fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
					value_over_all_average = append(value_over_all_average, *data4)
					fuzzy_self_points = fuzzy.SelfPoints
					fuzzy_work_points = fuzzy.WorkPoints
					fuzzy_co_worker_points = fuzzy.CoWorkerPoints
					value_over_all = fuzzy.ValueOverAll
					count = 1
					startDate = res[0]
				}
			}
			data1 := new(FuzzyAverage)
			data2 := new(FuzzyAverage)
			data3 := new(FuzzyAverage)
			data4 := new(FuzzyAverage)
			data1.Value = float32(fuzzy_self_points) / float32(count)
			data2.Value = float32(fuzzy_work_points) / float32(count)
			data3.Value = float32(fuzzy_co_worker_points) / float32(count)
			data4.Value = float32(value_over_all) / float32(count)
			data1.DateTime = startDate
			data2.DateTime = startDate
			data3.DateTime = startDate
			data4.DateTime = startDate
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)
			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	case "week":
		{
			dayName := ""
			if accountId != "" {
				rows, err := h.DB.Query(getHappinessScoreAllByAccountId, accountId)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				fuzzy := new(FuzzyValues)
				for rows.Next() {
					if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					currentTime, err := time.ParseInLocation(DDMMYYYYhhmmss, string(fuzzy.Time), location)
					if err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					res1 := strings.Split(string(fuzzy.Time), " ")
					day := currentTime.Weekday()
					if srtDate[0] == "" && dayName == "" {
						srtDate[0] = res1[0]
						dayName = day.String()
					}
					if dayName == day.String() {
						fuzzy_self_points += fuzzy.SelfPoints
						fuzzy_work_points += fuzzy.WorkPoints
						fuzzy_co_worker_points += fuzzy.CoWorkerPoints
						value_over_all += fuzzy.ValueOverAll
						srtDate[1] = res1[0]
						count++
					}
					if dayName != day.String() && dayName != "Friday" {
						dayName = day.String()
						fuzzy_self_points += fuzzy.SelfPoints
						fuzzy_work_points += fuzzy.WorkPoints
						fuzzy_co_worker_points += fuzzy.CoWorkerPoints
						value_over_all += fuzzy.ValueOverAll
						srtDate[1] = res1[0]
						count++
					}
					if dayName == "Friday" {
						srtDate[1] = res1[0]
						data1 := new(FuzzyAverage)
						data2 := new(FuzzyAverage)
						data3 := new(FuzzyAverage)
						data4 := new(FuzzyAverage)
						data1.Value = float32(fuzzy_self_points) / float32(count)
						data2.Value = float32(fuzzy_work_points) / float32(count)
						data3.Value = float32(fuzzy_co_worker_points) / float32(count)
						data4.Value = float32(value_over_all) / float32(count)
						data1.DateTime = srtDate[0] + " - " + srtDate[1]
						data2.DateTime = srtDate[0] + " - " + srtDate[1]
						data3.DateTime = srtDate[0] + " - " + srtDate[1]
						data4.DateTime = srtDate[0] + " - " + srtDate[1]
						fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
						fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
						fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
						value_over_all_average = append(value_over_all_average, *data4)
						fuzzy_self_points = 0
						fuzzy_work_points = 0
						fuzzy_co_worker_points = 0
						value_over_all = 0
						count = 0
						dayName = ""
						srtDate[0] = ""
					}

				}
				return c.JSON(http.StatusOK, echo.Map{
					"fuzzy_self_points_average":      fuzzy_self_points_average,
					"fuzzy_work_points_average":      fuzzy_work_points_average,
					"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
					"value_over_all_average":         value_over_all_average,
				})
			}
			rows, err := h.DB.Query(getHappinessScoreAll)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy := new(FuzzyValues)
			for rows.Next() {
				if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				currentTime, err := time.ParseInLocation(DDMMYYYYhhmmss, string(fuzzy.Time), location)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				res1 := strings.Split(string(fuzzy.Time), " ")
				day := currentTime.Weekday()
				if srtDate[0] == "" && dayName == "" {
					srtDate[0] = res1[0]
					dayName = day.String()
				}
				if dayName == day.String() {
					fuzzy_self_points += fuzzy.SelfPoints
					fuzzy_work_points += fuzzy.WorkPoints
					fuzzy_co_worker_points += fuzzy.CoWorkerPoints
					value_over_all += fuzzy.ValueOverAll
					srtDate[1] = res1[0]
					count++
				}
				if dayName != day.String() {
					dayName = day.String()
					fuzzy_self_points += fuzzy.SelfPoints
					fuzzy_work_points += fuzzy.WorkPoints
					fuzzy_co_worker_points += fuzzy.CoWorkerPoints
					value_over_all += fuzzy.ValueOverAll
					srtDate[1] = res1[0]
					count++
				}
				if dayName == "Friday" && fuzzy.AccountId == 300 {
					srtDate[1] = res1[0]
					data1 := new(FuzzyAverage)
					data2 := new(FuzzyAverage)
					data3 := new(FuzzyAverage)
					data4 := new(FuzzyAverage)
					data1.Value = float32(fuzzy_self_points) / float32(count)
					data2.Value = float32(fuzzy_work_points) / float32(count)
					data3.Value = float32(fuzzy_co_worker_points) / float32(count)
					data4.Value = float32(value_over_all) / float32(count)
					data1.DateTime = srtDate[0] + " - " + srtDate[1]
					data2.DateTime = srtDate[0] + " - " + srtDate[1]
					data3.DateTime = srtDate[0] + " - " + srtDate[1]
					data4.DateTime = srtDate[0] + " - " + srtDate[1]
					fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
					fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
					fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
					value_over_all_average = append(value_over_all_average, *data4)
					fuzzy_self_points = 0
					fuzzy_work_points = 0
					fuzzy_co_worker_points = 0
					value_over_all = 0
					count = 0
					dayName = ""
					srtDate[0] = ""
				}

			}
			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	case "month":
		{
			times := time.Now().In(location).Format(DDMMYYYYhhmmss)
			arr := strings.Split(times, "-")
			timeStart := "2023-" + arr[1] + "-01 00:00:00"
			timeStop := "2023-" + arr[1] + "-31 23:59:59"
			if accountId != "" {
				rows, err := h.DB.Query(getHappinessScoreAllByAccountIdAndDate, accountId, timeStart, timeStop)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				fuzzy := new(FuzzyValues)
				for rows.Next() {
					if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					res := strings.Split(string(fuzzy.Time), " ")
					if startDate == "" {
						startDate = res[0]
					}
					if startDate == res[0] {
						count++
						fuzzy_self_points += fuzzy.SelfPoints
						fuzzy_work_points += fuzzy.WorkPoints
						fuzzy_co_worker_points += fuzzy.CoWorkerPoints
						value_over_all += fuzzy.ValueOverAll
					}
					if startDate != res[0] {
						data1 := new(FuzzyAverage)
						data2 := new(FuzzyAverage)
						data3 := new(FuzzyAverage)
						data4 := new(FuzzyAverage)
						data1.Value = float32(fuzzy_self_points) / float32(count)
						data2.Value = float32(fuzzy_work_points) / float32(count)
						data3.Value = float32(fuzzy_co_worker_points) / float32(count)
						data4.Value = float32(value_over_all) / float32(count)
						data1.DateTime = startDate
						data2.DateTime = startDate
						data3.DateTime = startDate
						data4.DateTime = startDate
						fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
						fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
						fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
						value_over_all_average = append(value_over_all_average, *data4)
						fuzzy_self_points = fuzzy.SelfPoints
						fuzzy_work_points = fuzzy.WorkPoints
						fuzzy_co_worker_points = fuzzy.CoWorkerPoints
						value_over_all = fuzzy.ValueOverAll
						count = 1
						startDate = res[0]
					}
				}
				data1 := new(FuzzyAverage)
				data2 := new(FuzzyAverage)
				data3 := new(FuzzyAverage)
				data4 := new(FuzzyAverage)
				data1.Value = float32(fuzzy_self_points) / float32(count)
				data2.Value = float32(fuzzy_work_points) / float32(count)
				data3.Value = float32(fuzzy_co_worker_points) / float32(count)
				data4.Value = float32(value_over_all) / float32(count)
				data1.DateTime = startDate
				data2.DateTime = startDate
				data3.DateTime = startDate
				data4.DateTime = startDate
				fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
				fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
				fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
				value_over_all_average = append(value_over_all_average, *data4)
				return c.JSON(http.StatusOK, echo.Map{
					"fuzzy_self_points_average":      fuzzy_self_points_average,
					"fuzzy_work_points_average":      fuzzy_work_points_average,
					"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
					"value_over_all_average":         value_over_all_average,
				})
			}
			rows, err := h.DB.Query(getHappinessScoreAllByDate, timeStart, timeStop)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy := new(FuzzyValues)
			for rows.Next() {
				if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				res := strings.Split(string(fuzzy.Time), " ")
				if startDate == "" {
					startDate = res[0]
				}
				if startDate == res[0] {
					count++
					fuzzy_self_points += fuzzy.SelfPoints
					fuzzy_work_points += fuzzy.WorkPoints
					fuzzy_co_worker_points += fuzzy.CoWorkerPoints
					value_over_all += fuzzy.ValueOverAll
				}
				if startDate != res[0] {
					data1 := new(FuzzyAverage)
					data2 := new(FuzzyAverage)
					data3 := new(FuzzyAverage)
					data4 := new(FuzzyAverage)
					data1.Value = float32(fuzzy_self_points) / float32(count)
					data2.Value = float32(fuzzy_work_points) / float32(count)
					data3.Value = float32(fuzzy_co_worker_points) / float32(count)
					data4.Value = float32(value_over_all) / float32(count)
					data1.DateTime = startDate
					data2.DateTime = startDate
					data3.DateTime = startDate
					data4.DateTime = startDate
					fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
					fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
					fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
					value_over_all_average = append(value_over_all_average, *data4)
					fuzzy_self_points = fuzzy.SelfPoints
					fuzzy_work_points = fuzzy.WorkPoints
					fuzzy_co_worker_points = fuzzy.CoWorkerPoints
					value_over_all = fuzzy.ValueOverAll
					count = 1
					startDate = res[0]
				}
			}
			data1 := new(FuzzyAverage)
			data2 := new(FuzzyAverage)
			data3 := new(FuzzyAverage)
			data4 := new(FuzzyAverage)
			data1.Value = float32(fuzzy_self_points) / float32(count)
			data2.Value = float32(fuzzy_work_points) / float32(count)
			data3.Value = float32(fuzzy_co_worker_points) / float32(count)
			data4.Value = float32(value_over_all) / float32(count)
			data1.DateTime = startDate
			data2.DateTime = startDate
			data3.DateTime = startDate
			data4.DateTime = startDate
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)
			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	case "year":
		{
			times := time.Now().In(location).Format(DDMMYYYYhhmmss)
			arr := strings.Split(times, "-")
			timeStart := arr[0] + "-" + "01-01 00:00:00"
			timeStop := arr[0] + "-" + "12-31 23:59:59"
			if accountId != "" {
				rows, err := h.DB.Query(getHappinessScoreAllByAccountIdAndDate, accountId, timeStart, timeStop)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				fuzzy := new(FuzzyValues)
				for rows.Next() {
					if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
						return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
					}
					res := strings.Split(string(fuzzy.Time), " ")
					if startDate == "" {
						startDate = res[0]
					}
					if startDate == res[0] {
						count++
						fuzzy_self_points += fuzzy.SelfPoints
						fuzzy_work_points += fuzzy.WorkPoints
						fuzzy_co_worker_points += fuzzy.CoWorkerPoints
						value_over_all += fuzzy.ValueOverAll
					}
					if startDate != res[0] {
						data1 := new(FuzzyAverage)
						data2 := new(FuzzyAverage)
						data3 := new(FuzzyAverage)
						data4 := new(FuzzyAverage)
						data1.Value = float32(fuzzy_self_points) / float32(count)
						data2.Value = float32(fuzzy_work_points) / float32(count)
						data3.Value = float32(fuzzy_co_worker_points) / float32(count)
						data4.Value = float32(value_over_all) / float32(count)
						data1.DateTime = startDate
						data2.DateTime = startDate
						data3.DateTime = startDate
						data4.DateTime = startDate
						fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
						fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
						fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
						value_over_all_average = append(value_over_all_average, *data4)
						fuzzy_self_points = fuzzy.SelfPoints
						fuzzy_work_points = fuzzy.WorkPoints
						fuzzy_co_worker_points = fuzzy.CoWorkerPoints
						value_over_all = fuzzy.ValueOverAll
						count = 1
						startDate = res[0]
					}
				}
				data1 := new(FuzzyAverage)
				data2 := new(FuzzyAverage)
				data3 := new(FuzzyAverage)
				data4 := new(FuzzyAverage)
				data1.Value = float32(fuzzy_self_points) / float32(count)
				data2.Value = float32(fuzzy_work_points) / float32(count)
				data3.Value = float32(fuzzy_co_worker_points) / float32(count)
				data4.Value = float32(value_over_all) / float32(count)
				data1.DateTime = startDate
				data2.DateTime = startDate
				data3.DateTime = startDate
				data4.DateTime = startDate
				fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
				fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
				fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
				value_over_all_average = append(value_over_all_average, *data4)
				return c.JSON(http.StatusOK, echo.Map{
					"fuzzy_self_points_average":      fuzzy_self_points_average,
					"fuzzy_work_points_average":      fuzzy_work_points_average,
					"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
					"value_over_all_average":         value_over_all_average,
				})
			}
			rows, err := h.DB.Query(getHappinessScoreAllByDate, timeStart, timeStop)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy := new(FuzzyValues)
			for rows.Next() {
				if err := rows.Scan(&fuzzy.Id, &fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints, &fuzzy.ValueOverAll, &fuzzy.Time, &fuzzy.AccountId); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				res := strings.Split(string(fuzzy.Time), " ")
				if startDate == "" {
					startDate = res[0]
				}
				if startDate == res[0] {
					count++
					fuzzy_self_points += fuzzy.SelfPoints
					fuzzy_work_points += fuzzy.WorkPoints
					fuzzy_co_worker_points += fuzzy.CoWorkerPoints
					value_over_all += fuzzy.ValueOverAll
				}
				if startDate != res[0] {
					data1 := new(FuzzyAverage)
					data2 := new(FuzzyAverage)
					data3 := new(FuzzyAverage)
					data4 := new(FuzzyAverage)
					data1.Value = float32(fuzzy_self_points) / float32(count)
					data2.Value = float32(fuzzy_work_points) / float32(count)
					data3.Value = float32(fuzzy_co_worker_points) / float32(count)
					data4.Value = float32(value_over_all) / float32(count)
					data1.DateTime = startDate
					data2.DateTime = startDate
					data3.DateTime = startDate
					data4.DateTime = startDate
					fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
					fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
					fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
					value_over_all_average = append(value_over_all_average, *data4)
					fuzzy_self_points = fuzzy.SelfPoints
					fuzzy_work_points = fuzzy.WorkPoints
					fuzzy_co_worker_points = fuzzy.CoWorkerPoints
					value_over_all = fuzzy.ValueOverAll
					count = 1
					startDate = res[0]
				}
			}
			data1 := new(FuzzyAverage)
			data2 := new(FuzzyAverage)
			data3 := new(FuzzyAverage)
			data4 := new(FuzzyAverage)
			data1.Value = float32(fuzzy_self_points) / float32(count)
			data2.Value = float32(fuzzy_work_points) / float32(count)
			data3.Value = float32(fuzzy_co_worker_points) / float32(count)
			data4.Value = float32(value_over_all) / float32(count)
			data1.DateTime = startDate
			data2.DateTime = startDate
			data3.DateTime = startDate
			data4.DateTime = startDate
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)
			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	default:
		return c.JSON(http.StatusBadRequest, "")
	}
}

func (h *Handler) GetHappinessByUserId(c echo.Context) error {
	CheckHTTP()
	listHappiness := new(ResponseGetHappines)
	happiness := new(models.HappinessPoint)
	hpPoint := new(Record)
	userId := c.Param("id")
	period := c.Param("period")
	if period != ":period" {
		startDate := ""
		stopDate := ""
		if period == "week" {
			startDate = time.Now().UTC().Format(YYYYMMDD)
			stopDate = time.Now().Add(time.Hour * -168).UTC().Format(YYYYMMDD)
		} else if period == "month" {
			startDate = time.Now().UTC().Format(YYYYMMDD)
			stopDate = time.Now().Add(time.Hour * -720).UTC().Format(YYYYMMDD)
		} else if period == "day" {
			startDate = time.Now().UTC().Format(YYYYMMDD)
			stopDate = time.Now().UTC().Format(YYYYMMDD)
		}
		rows, err := h.DB.Query(getHappinessByUserIdAndDate, userId, startDate, stopDate)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		listHappiness.Id, _ = strconv.Atoi(userId)
		listHappiness.Period = period
		for rows.Next() {
			if err := rows.Scan(&happiness.Id, &happiness.AccountId, &happiness.Selfpoints,
				&happiness.Workpoints, &happiness.Copoints, &happiness.TimeStamp); err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy, errf := FuzzyCalculatorAll(happiness.Selfpoints, happiness.Workpoints, happiness.Copoints)
			if errf != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			hpPoint.Hppoints.SelfPoints = happiness.Selfpoints
			hpPoint.Hppoints.WorkPoints = happiness.Workpoints
			hpPoint.Hppoints.CoWorkerPoints = happiness.Copoints
			hpPoint.Hppoints.FuzzyValue = fuzzy.Value
			hpPoint.Date = string(happiness.TimeStamp)
			listHappiness.Records = append(listHappiness.Records, *hpPoint)
		}
	} else if period == ":period" {
		rows, err := h.DB.Query(getHappinessByUserId, userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		listHappiness.Id, _ = strconv.Atoi(userId)
		listHappiness.Period = ""
		for rows.Next() {
			if err := rows.Scan(&happiness.Id, &happiness.AccountId, &happiness.Selfpoints,
				&happiness.Workpoints, &happiness.Copoints, &happiness.TimeStamp); err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			fuzzy, errf := FuzzyCalculatorAll(happiness.Selfpoints, happiness.Workpoints, happiness.Copoints)
			if errf != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			hpPoint.Hppoints.SelfPoints = happiness.Selfpoints
			hpPoint.Hppoints.WorkPoints = happiness.Workpoints
			hpPoint.Hppoints.CoWorkerPoints = happiness.Copoints
			hpPoint.Hppoints.FuzzyValue = fuzzy.Value
			hpPoint.Date = string(happiness.TimeStamp)
			listHappiness.Records = append(listHappiness.Records, *hpPoint)
		}
	}
	if len(listHappiness.Records) != 0 {
		return c.JSON(http.StatusOK, listHappiness)
	}
	return c.String(http.StatusNoContent, "")
}
func (h *Handler) GetHappinessScorePositionAllTime(c echo.Context) error {
	fuzzy := new(FuzzyValuesByPoistion)
	var fuzzy_self_points_average []FuzzyAverageByPoistion
	var fuzzy_work_points_average []FuzzyAverageByPoistion
	var fuzzy_co_worker_points_average []FuzzyAverageByPoistion
	var value_over_all_average []FuzzyAverageByPoistion
	departmentId := c.QueryParam("department_id")
	positionId := 1
	fuzzy_self_points := 0
	fuzzy_work_points := 0
	fuzzy_co_worker_points := 0
	value_over_all := 0
	count := 0
	positionNane := ""
	if departmentId == "" {
		return c.String(http.StatusBadRequest, "department_id is null")
	}
	rows, err := h.DB.Query(getHappinessScorePosition, departmentId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	for rows.Next() {
		if err := rows.Scan(&fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints,
			&fuzzy.ValueOverAll, &fuzzy.AccountId, &fuzzy.PositionId, &fuzzy.PositionName); err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		if positionNane == "" {
			positionNane = fuzzy.PositionName
		}
		if positionId == fuzzy.PositionId {
			fuzzy_self_points += fuzzy.SelfPoints
			fuzzy_work_points += fuzzy.WorkPoints
			fuzzy_co_worker_points += fuzzy.CoWorkerPoints
			value_over_all += fuzzy.ValueOverAll
			count++
		}
		if positionId != fuzzy.PositionId {
			data1 := new(FuzzyAverageByPoistion)
			data2 := new(FuzzyAverageByPoistion)
			data3 := new(FuzzyAverageByPoistion)
			data4 := new(FuzzyAverageByPoistion)
			data1.Value = float32(fuzzy_self_points) / float32(count)
			data2.Value = float32(fuzzy_work_points) / float32(count)
			data3.Value = float32(fuzzy_co_worker_points) / float32(count)
			data4.Value = float32(value_over_all) / float32(count)
			data1.Position = positionNane
			data2.Position = positionNane
			data3.Position = positionNane
			data4.Position = positionNane
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)
			fuzzy_self_points = fuzzy.SelfPoints
			fuzzy_work_points = fuzzy.WorkPoints
			fuzzy_co_worker_points = fuzzy.CoWorkerPoints
			value_over_all = fuzzy.ValueOverAll
			count = 1
			positionNane = fuzzy.PositionName
			positionId = fuzzy.PositionId
		}
	}
	data1 := new(FuzzyAverageByPoistion)
	data2 := new(FuzzyAverageByPoistion)
	data3 := new(FuzzyAverageByPoistion)
	data4 := new(FuzzyAverageByPoistion)
	data1.Value = float32(fuzzy_self_points) / float32(count)
	data2.Value = float32(fuzzy_work_points) / float32(count)
	data3.Value = float32(fuzzy_co_worker_points) / float32(count)
	data4.Value = float32(value_over_all) / float32(count)
	data1.Position = positionNane
	data2.Position = positionNane
	data3.Position = positionNane
	data4.Position = positionNane
	fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
	fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
	fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
	value_over_all_average = append(value_over_all_average, *data4)

	return c.JSON(http.StatusOK, echo.Map{
		"fuzzy_self_points_average":      fuzzy_self_points_average,
		"fuzzy_work_points_average":      fuzzy_work_points_average,
		"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
		"value_over_all_average":         value_over_all_average,
	})
}
func (h *Handler) GetHappinessScoreDepartMentAllTime(c echo.Context) error {
	fuzzy := new(FuzzyValuesByDepartMent)
	var fuzzy_self_points_average []FuzzyAverageByDepartMent
	var fuzzy_work_points_average []FuzzyAverageByDepartMent
	var fuzzy_co_worker_points_average []FuzzyAverageByDepartMent
	var value_over_all_average []FuzzyAverageByDepartMent
	departmentId := 1
	fuzzy_self_points_ce := 0
	fuzzy_work_points_ce := 0
	fuzzy_co_worker_points_ce := 0
	value_over_all_ce := 0
	fuzzy_self_points_mle := 0
	fuzzy_work_points_mle := 0
	fuzzy_co_worker_points_mle := 0
	value_over_all_mle := 0
	fuzzy_self_points_it := 0
	fuzzy_work_points_it := 0
	fuzzy_co_worker_points_it := 0
	value_over_all_it := 0
	count_ce := 0
	count_mle := 0
	count_it := 0
	departMentNane := ""
	switch period := c.QueryParam("period"); period {
	case "":
		{
			fuzzy_self_points := 0
			fuzzy_work_points := 0
			fuzzy_co_worker_points := 0
			value_over_all := 0
			count := 0
			rows, err := h.DB.Query(getHappinessScoreDepartment)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			for rows.Next() {
				if err := rows.Scan(&fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints,
					&fuzzy.ValueOverAll, &fuzzy.AccountId, &fuzzy.DepartMentId, &fuzzy.DepartMentName); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				if departMentNane == "" {
					departMentNane = fuzzy.DepartMentName
				}
				if departmentId == fuzzy.DepartMentId {
					fuzzy_self_points += fuzzy.SelfPoints
					fuzzy_work_points += fuzzy.WorkPoints
					fuzzy_co_worker_points += fuzzy.CoWorkerPoints
					value_over_all += fuzzy.ValueOverAll
					count++
				}
				if departmentId != fuzzy.DepartMentId {
					data1 := new(FuzzyAverageByDepartMent)
					data2 := new(FuzzyAverageByDepartMent)
					data3 := new(FuzzyAverageByDepartMent)
					data4 := new(FuzzyAverageByDepartMent)
					data1.Value = float32(fuzzy_self_points) / float32(count)
					data2.Value = float32(fuzzy_work_points) / float32(count)
					data3.Value = float32(fuzzy_co_worker_points) / float32(count)
					data4.Value = float32(value_over_all) / float32(count)
					data1.Department = departMentNane
					data2.Department = departMentNane
					data3.Department = departMentNane
					data4.Department = departMentNane
					fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
					fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
					fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
					value_over_all_average = append(value_over_all_average, *data4)
					fuzzy_self_points = fuzzy.SelfPoints
					fuzzy_work_points = fuzzy.WorkPoints
					fuzzy_co_worker_points = fuzzy.CoWorkerPoints
					value_over_all = fuzzy.ValueOverAll
					count = 1
					departMentNane = fuzzy.DepartMentName
					departmentId = fuzzy.DepartMentId
				}
			}
			data1 := new(FuzzyAverageByDepartMent)
			data2 := new(FuzzyAverageByDepartMent)
			data3 := new(FuzzyAverageByDepartMent)
			data4 := new(FuzzyAverageByDepartMent)
			data1.Value = float32(fuzzy_self_points) / float32(count)
			data2.Value = float32(fuzzy_work_points) / float32(count)
			data3.Value = float32(fuzzy_co_worker_points) / float32(count)
			data4.Value = float32(value_over_all) / float32(count)
			data1.Department = departMentNane
			data2.Department = departMentNane
			data3.Department = departMentNane
			data4.Department = departMentNane
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)

			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	case "month":
		{
			location := time.FixedZone("UTC+7", 7*60*60)
			times := time.Now().In(location).Format(DDMMYYYYhhmmss)
			arr := strings.Split(times, "-")
			timeStart := "2023-" + arr[1] + "-01 00:00:00"
			timeStop := "2023-" + arr[1] + "-31 23:59:59"
			rows, err := h.DB.Query(getHappinessScoreDepartmentByDate, timeStart, timeStop)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			for rows.Next() {
				if err := rows.Scan(&fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints,
					&fuzzy.ValueOverAll, &fuzzy.AccountId, &fuzzy.DepartMentId, &fuzzy.DepartMentName); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				if fuzzy.DepartMentName == "CE" {
					fuzzy_self_points_ce += fuzzy.SelfPoints
					fuzzy_work_points_ce += fuzzy.WorkPoints
					fuzzy_co_worker_points_ce += fuzzy.CoWorkerPoints
					value_over_all_ce += fuzzy.ValueOverAll
					count_ce++
				}
				if fuzzy.DepartMentName == "IT" {
					fuzzy_self_points_it += fuzzy.SelfPoints
					fuzzy_work_points_it += fuzzy.WorkPoints
					fuzzy_co_worker_points_it += fuzzy.CoWorkerPoints
					value_over_all_it += fuzzy.ValueOverAll
					count_it++
				}
				if fuzzy.DepartMentName == "MLE" {
					fuzzy_self_points_mle += fuzzy.SelfPoints
					fuzzy_work_points_mle += fuzzy.WorkPoints
					fuzzy_co_worker_points_mle += fuzzy.CoWorkerPoints
					value_over_all_mle += fuzzy.ValueOverAll
					count_mle++
				}
			}

			data1 := new(FuzzyAverageByDepartMent)
			data2 := new(FuzzyAverageByDepartMent)
			data3 := new(FuzzyAverageByDepartMent)
			data4 := new(FuzzyAverageByDepartMent)
			data1.Value = float32(fuzzy_self_points_ce) / float32(count_ce)
			data2.Value = float32(fuzzy_work_points_ce) / float32(count_ce)
			data3.Value = float32(fuzzy_co_worker_points_ce) / float32(count_ce)
			data4.Value = float32(value_over_all_ce) / float32(count_ce)
			data1.Department = "CE"
			data2.Department = "CE"
			data3.Department = "CE"
			data4.Department = "CE"
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)

			data1.Value = float32(fuzzy_self_points_it) / float32(count_it)
			data2.Value = float32(fuzzy_work_points_it) / float32(count_it)
			data3.Value = float32(fuzzy_co_worker_points_it) / float32(count_it)
			data4.Value = float32(value_over_all_it) / float32(count_it)
			data1.Department = "IT"
			data2.Department = "IT"
			data3.Department = "IT"
			data4.Department = "IT"
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)

			data1.Value = float32(fuzzy_self_points_mle) / float32(count_mle)
			data2.Value = float32(fuzzy_work_points_mle) / float32(count_mle)
			data3.Value = float32(fuzzy_co_worker_points_mle) / float32(count_mle)
			data4.Value = float32(value_over_all_mle) / float32(count_mle)
			data1.Department = "MLE"
			data2.Department = "MLE"
			data3.Department = "MLE"
			data4.Department = "MLE"
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)
			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	case "year":
		{
			location := time.FixedZone("UTC+7", 7*60*60)
			times := time.Now().In(location).Format(DDMMYYYYhhmmss)
			arr := strings.Split(times, "-")
			timeStart := arr[0] + "-" + "01-01 00:00:00"
			timeStop := arr[0] + "-" + "12-31 23:59:59"
			rows, err := h.DB.Query(getHappinessScoreDepartmentByDate, timeStart, timeStop)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
			}
			for rows.Next() {
				if err := rows.Scan(&fuzzy.SelfPoints, &fuzzy.WorkPoints, &fuzzy.CoWorkerPoints,
					&fuzzy.ValueOverAll, &fuzzy.AccountId, &fuzzy.DepartMentId, &fuzzy.DepartMentName); err != nil {
					return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
				}
				if fuzzy.DepartMentName == "CE" {
					fuzzy_self_points_ce += fuzzy.SelfPoints
					fuzzy_work_points_ce += fuzzy.WorkPoints
					fuzzy_co_worker_points_ce += fuzzy.CoWorkerPoints
					value_over_all_ce += fuzzy.ValueOverAll
					count_ce++
				}
				if fuzzy.DepartMentName == "IT" {
					fuzzy_self_points_it += fuzzy.SelfPoints
					fuzzy_work_points_it += fuzzy.WorkPoints
					fuzzy_co_worker_points_it += fuzzy.CoWorkerPoints
					value_over_all_it += fuzzy.ValueOverAll
					count_it++
				}
				if fuzzy.DepartMentName == "MLE" {
					fuzzy_self_points_mle += fuzzy.SelfPoints
					fuzzy_work_points_mle += fuzzy.WorkPoints
					fuzzy_co_worker_points_mle += fuzzy.CoWorkerPoints
					value_over_all_mle += fuzzy.ValueOverAll
					count_mle++
				}
			}

			data1 := new(FuzzyAverageByDepartMent)
			data2 := new(FuzzyAverageByDepartMent)
			data3 := new(FuzzyAverageByDepartMent)
			data4 := new(FuzzyAverageByDepartMent)
			data1.Value = float32(fuzzy_self_points_ce) / float32(count_ce)
			data2.Value = float32(fuzzy_work_points_ce) / float32(count_ce)
			data3.Value = float32(fuzzy_co_worker_points_ce) / float32(count_ce)
			data4.Value = float32(value_over_all_ce) / float32(count_ce)
			data1.Department = "CE"
			data2.Department = "CE"
			data3.Department = "CE"
			data4.Department = "CE"
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)

			data1.Value = float32(fuzzy_self_points_it) / float32(count_it)
			data2.Value = float32(fuzzy_work_points_it) / float32(count_it)
			data3.Value = float32(fuzzy_co_worker_points_it) / float32(count_it)
			data4.Value = float32(value_over_all_it) / float32(count_it)
			data1.Department = "IT"
			data2.Department = "IT"
			data3.Department = "IT"
			data4.Department = "IT"
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)

			data1.Value = float32(fuzzy_self_points_mle) / float32(count_mle)
			data2.Value = float32(fuzzy_work_points_mle) / float32(count_mle)
			data3.Value = float32(fuzzy_co_worker_points_mle) / float32(count_mle)
			data4.Value = float32(value_over_all_mle) / float32(count_mle)
			data1.Department = "MLE"
			data2.Department = "MLE"
			data3.Department = "MLE"
			data4.Department = "MLE"
			fuzzy_self_points_average = append(fuzzy_self_points_average, *data1)
			fuzzy_work_points_average = append(fuzzy_work_points_average, *data2)
			fuzzy_co_worker_points_average = append(fuzzy_co_worker_points_average, *data3)
			value_over_all_average = append(value_over_all_average, *data4)
			return c.JSON(http.StatusOK, echo.Map{
				"fuzzy_self_points_average":      fuzzy_self_points_average,
				"fuzzy_work_points_average":      fuzzy_work_points_average,
				"fuzzy_co_worker_points_average": fuzzy_co_worker_points_average,
				"value_over_all_average":         value_over_all_average,
			})
		}
	default:
		{
			return c.JSON(http.StatusBadRequest, "")
		}

	}
}
