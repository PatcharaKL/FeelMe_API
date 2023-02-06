package tokens

import (
	"database/sql"
	"strings"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	Signingkey               = "GVebOWpKrqyZ9RwPXzazpNpcmA6njskh"
	getAccountById           = "SELECT * FROM accounts WHERE account_id=?"
	getRefresh               = "SELECT * FROM refresh_token WHERE refreshToken = ? && isValid = ? && exp >= ?"
	createRefreshToken       = `INSERT INTO refresh_token (refreshToken, account_id, exp, isValid) VALUES (?, ?, ?, ?)RETURNING refreshToken;`
	updateStatusRefreshToken = "UPDATE refresh_token SET  isValid = ? WHERE refreshToken = ?"
)

type Refreshtoken struct {
	Refreshtoken string `json:"refreshToken" form:"refreshToken" query:"refreshToken"`
}
type Token struct {
	Refreshtokens string  `json:"refreshToken"  form:"refreshToken"  query:"refreshToken"`
	AccountId     int     `json:"accountId" form:"accountId" query:"accountId"`
	Exp           []uint8 `json:"exp" form:"exp" query:"exp"`
	IsValid       bool    `json:"isValid" form:"isValid" query:"isValid"`
}
type JwtCustomClaims struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Role         int    `json:"role"`
	AccountId    int    `json:"accountId"`
	DepartmentId int    `json:"departmentId"`
	CompanyId    int    `json:"companyId"`
	jwt.RegisteredClaims
}
type Handler struct {
	DB *sql.DB
}

func NewApplication(db *sql.DB) *Handler {
	return &Handler{db}
}

type Err struct {
	Message string `json:"message"`
}

func GeneraterRefreshToken() string {
	refreshToken := []string{}
	for i := 0; i < 4; i++ {
		id := uuid.New()
		refreshToken = append(refreshToken, id.String())
	}
	return strings.Join(refreshToken, "-")
}
func GeneraterTokenAccess(ac models.Account) (string, error) {
	claims := &JwtCustomClaims{ac.Email, ac.Name, ac.Surname, ac.PositionId, ac.AccountId, ac.DepartmentId,
		ac.CompanyId, jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(Signingkey))
	if err != nil {
		return "", err
	}
	return t, nil
}
