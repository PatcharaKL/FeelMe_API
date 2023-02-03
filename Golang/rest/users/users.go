package users

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
)

const (
	getAccountByEmail          = "SELECT * FROM accounts WHERE email=?"
	getAccountById             = "SELECT * FROM accounts WHERE account_id=?"
	getRefreshTokenByAccountId = "SELECT refreshToken FROM refresh_token WHERE account_id=? && isValid=?"
	createRefreshToken         = `INSERT INTO refresh_token (refreshToken, account_id, exp, isValid) VALUES (?, ?, ?, ?)RETURNING refreshToken;`
	updateStatusRefreshToken   = "UPDATE refresh_token SET  isValid = ? WHERE refreshToken = ?"
)

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
type userLogin struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type userLoginDTO struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

func NewApplication(db *sql.DB) *Handler {
	return &Handler{db}
}

type Err struct {
	Message string `json:"message"`
}
