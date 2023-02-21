package users

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	getAccountByEmail           = "SELECT * FROM accounts WHERE email=?"
	getHappinessByUserId        = "SELECT * FROM deily_happiness_points WHERE  account_id=?"
	getHappinessByUserIdAndDate = "SELECT * FROM deily_happiness_points WHERE  account_id=? && timestamp <= ? && timestamp >= ?"
	getAccountById              = "SELECT * FROM accounts WHERE account_id=?"
	getRefreshTokenByAccountId  = "SELECT refreshToken FROM refresh_token WHERE account_id=? && isValid=?"
	createRefreshToken          = `INSERT INTO refresh_token (refreshToken, account_id, exp, isValid) VALUES (?, ?, ?, ?)RETURNING refreshToken;`
	updateStatusRefreshToken    = "UPDATE refresh_token SET  isValid = ? WHERE refreshToken = ?"
	createdHappinessPoint       = "INSERT INTO deily_happiness_points (account_id,seif_point,work_point,co_worker_point,timestamp) VALUES (?, ?, ?, ?,?)RETURNING id;"
)

type HapPointRequest struct {
	Selfpoints int `json:"seif_points"`
	Workpoints int `json:"work_points"`
	Copoints   int `json:"co_worker_points"`
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
