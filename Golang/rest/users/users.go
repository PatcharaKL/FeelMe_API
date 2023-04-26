package users

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	getAccountByEmail           = "SELECT * FROM accounts WHERE email=?"
	getHappinessByUserId        = "SELECT * FROM deily_happiness_points WHERE  account_id=?;"
	getUserFullNameByUserId     = "SELECT name,surname FROM feelme_db.accounts WHERE account_id=? ;"
	getUserSearchByName         = "SELECT account_id,name,surname,hp,level,avatar_url,positions.position_name,departments.department_name,companies.company_name FROM feelme_db.accounts join feelme_db.positions ON feelme_db.positions.position_id = feelme_db.accounts.position_id join feelme_db.departments ON feelme_db.departments.department_id = feelme_db.accounts.department_id join feelme_db.companies ON feelme_db.companies.company_id = feelme_db.accounts.company_id where name like ?;"
	getHappinessByUserIdAndDate = "SELECT * FROM deily_happiness_points WHERE  account_id=? && timestamp <= ? && timestamp >= ?"
	getUserDetail               = "SELECT account_id,name,surname,hp,level,avatar_url,positions.position_name,departments.department_name,companies.company_name FROM feelme_db.accounts join feelme_db.positions ON feelme_db.positions.position_id = feelme_db.accounts.position_id join feelme_db.departments ON feelme_db.departments.department_id = feelme_db.accounts.department_id join feelme_db.companies ON feelme_db.companies.company_id = feelme_db.accounts.company_id;"
	GetUserByUserId             = "SELECT account_id,name,surname,hp,level,avatar_url,positions.position_name,departments.department_name,companies.company_name FROM feelme_db.accounts join feelme_db.positions ON feelme_db.positions.position_id = feelme_db.accounts.position_id join feelme_db.departments ON feelme_db.departments.department_id = feelme_db.accounts.department_id join feelme_db.companies ON feelme_db.companies.company_id = feelme_db.accounts.company_id WHERE account_id=? ;"
	getRefreshTokenByAccountId  = "SELECT refreshToken FROM refresh_token WHERE account_id=? && isValid=?"
	createRefreshToken          = `INSERT INTO refresh_token (refreshToken, account_id, exp, isValid) VALUES (?, ?, ?, ?)RETURNING refreshToken;`
	updateStatusRefreshToken    = "UPDATE refresh_token SET  isValid = ? WHERE refreshToken = ?"
	UpdateUserData              = "UPDATE accounts SET  hp = ? WHERE account_id = ?"
	createdHappinessPoint       = "INSERT INTO deily_happiness_points (account_id,seif_point,work_point,co_worker_point,timestamp) VALUES (?, ?, ?, ?,?)RETURNING id;"
	createdLogTimeStamp         = "INSERT INTO feelme_db.log_timestamps (username,timestamp_type,user_id,time) VALUES (?, ?, ?, ?) RETURNING id;"
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
