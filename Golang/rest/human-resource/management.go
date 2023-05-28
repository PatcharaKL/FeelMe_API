package humanresource

import (
	"net/http"
	"time"

	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
	"github.com/PatcharaKL/FeelMe_API/service"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type EditUser struct {
	AccountId    int    `json:"account_id"`
	Name         string `json:"name"`
	SurName      string `json:"surname"`
	PositionId   int    `json:"position_id"`
	DepartmentId int    `json:"department_id"`
}

type TimeStamp struct {
	NameUser string  `json:"nameUser"`
	LogType  string  `json:"logType"`
	Time     []uint8 `json:"timestamp"`
}
type TimeStampData struct {
	NameUser string `json:"nameUser"`
	LogType  string `json:"logType"`
	Time     string `json:"timestamp"`
}
type CreatedEmployee struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	SurName      string `json:"surname"`
	Password     string `json:"password"`
	PositionId   int    `json:"position_id"`
	DepartmentId int    `json:"department_id"`
	CompanyId    int    `json:"company_id"`
}

func HashPassword(password []byte) (string, error) {
	cost := 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", nil
	}
	return string(hashedPassword), nil
}

const (
	YYYYMMDD       = "2006-01-02"
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func (h *Handler) EditProfileEmployee(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	editUser := new(EditUser)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	if claims.Role != 4 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Status":  false,
			"Message": "You Not HR",
		})
	}
	if err := c.Bind(editUser); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	stmt, err := h.DB.Prepare(editUserProfile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	if _, err := stmt.Exec(editUser.Name, editUser.SurName, editUser.PositionId, editUser.DepartmentId, editUser.AccountId); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Status":  true,
		"Message": "Succeed",
	})
}
func (h *Handler) UpdateUserImageProfile(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	if claims.Role != 4 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Status":  false,
			"Message": "You Not HR",
		})
	}
	accountId := c.QueryParam("account_id")
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	fileName := file.Filename
	uploadFile, uploadErr := service.UploadService("feelme-image/profile", fileName, src)
	if uploadErr != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: uploadErr.Error()})
	}
	stmt, err := h.DB.Prepare(UpdateProfileImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if _, err := stmt.Exec(uploadFile, accountId); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Status":    true,
		"Message":   "Success",
		"AccountID": accountId,
	})
}

func (h *Handler) CreatedUser(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	employee := new(CreatedEmployee)
	claims := user.Claims.(*tokens.JwtCustomClaims)
	location := time.FixedZone("UTC+7", 7*60*60)

	times := time.Now().In(location).Format(DDMMYYYYhhmmss)
	if claims.Role != 4 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Status":  false,
			"Message": "You Not HR",
		})
	}
	if err := c.Bind(employee); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	password := []byte(employee.Password)
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	id := 0
	if err := h.DB.QueryRow(createdEmployee, employee.Email, hashedPassword, employee.Name, employee.SurName,
		"https://cdn-icons-png.flaticon.com/512/236/236831.png", times, true, 100, 1, times,
		employee.DepartmentId, employee.PositionId, employee.CompanyId).Scan(&id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Status":  true,
		"Message": "Succeed",
	})
}
