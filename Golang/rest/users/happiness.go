package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/labstack/echo/v4"
)

const YYYYMMDD = "2006-01-02"

func (h *Handler) HappinesspointHandler(c echo.Context) error {

	userId := c.Param("id")

	hpyPointBody := new(HapPointRequest)
	if err := c.Bind(hpyPointBody); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	hpyPoint := models.HappinessPoint{}
	if err := h.DB.QueryRow(createdHappinessPoint, userId, hpyPointBody.Selfpoints, hpyPointBody.Workpoints, hpyPointBody.Copoints, time.Now()).Scan(&hpyPoint.Id); err != nil {

		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"user_id":          userId,
		"self_points":      hpyPointBody.Selfpoints,
		"work_points":      hpyPointBody.Workpoints,
		"co_worker_points": hpyPointBody.Copoints,
	})

}

type Hppoint struct {
	SelfPoints     int     `json:"self_points"`
	WorkPoints     int     `json:"work_points"`
	CoWorkerPoints int     `json:"co_worker_points"`
	FuzzyValue     float64 `json:"fuzzy_value"`
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

func FuzzyCalculator(self_points int, work_points int, co_points int) (*Value, error) {
	http_name := fmt.Sprintf("http://127.0.0.1:8000/v1/fuzzy?self_hp=%d&work_hp=%d&co_worker_hp=%d", self_points, work_points, co_points)
	vauel := new(Value)
	client := &http.Client{}
	req, err := http.NewRequest("GET", http_name, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	json.NewDecoder(resp.Body).Decode(vauel)
	return vauel, nil
}

func (h *Handler) GetHappinessByUserId(c echo.Context) error {
	listHappiness := new(ResponseGetHappines)
	happiness := new(models.HappinessPoint)
	hpPoint := new(Record)
	userId := c.Param("id")
	period := c.Param("period")
	if period != ":period" {
		startDate := ""
		stopDate := ""
		if period == "weeky" {
			startDate = time.Now().UTC().Format(YYYYMMDD)
			stopDate = time.Now().Add(time.Hour * -168).UTC().Format(YYYYMMDD)
		}
		if period == "month" {
			startDate = time.Now().UTC().Format(YYYYMMDD)
			stopDate = time.Now().Add(time.Hour * -720).UTC().Format(YYYYMMDD)
		}
		if period == "day" {
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
			fuzzy, err := FuzzyCalculator(happiness.Selfpoints, happiness.Workpoints, happiness.Copoints)
			if err != nil {
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
			fuzzy, err := FuzzyCalculator(happiness.Selfpoints, happiness.Workpoints, happiness.Copoints)
			if err != nil {
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
