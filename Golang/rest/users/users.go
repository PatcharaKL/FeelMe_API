package users

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	getAccountById             = "SELECT * FROM accounts WHERE email=?"
	getRefreshTokenByAccountId = "SELECT refreshToken FROM refresh_token WHERE account_id=? && isValid=?"
	createRefreshToken         = `INSERT INTO refresh_token (refreshToken, account_id, exp, isValid) VALUES (?, ?, ?, ?)RETURNING refreshToken;`
	updateStatusRefreshToken   = "UPDATE refresh_token SET  isValid = ? WHERE refreshToken = ?"
)

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

type Account struct {
	AccountId    int            `json:"account_id" query:"account_id"`
	Email        string         `json:"email" query:"email"`
	PasswordHash string         `json:"password_hash" query:"password_hash"`
	Name         string         `json:"name" query:"name"`
	Surname      string         `json:"surname" query:"surname"`
	AvatarUrl    sql.NullString `json:"avatar_url" query:"avatar_url"`
	ApplyDate    []uint8        `json:"apply_date" query:"apply_date"`
	IsActive     bool           `json:"is_active" query:"is_active"`
	Hp           int            `json:"hp" query:"hp"`
	Level        int            `json:"level" query:"level"`
	Created      []uint8        `json:"created" query:"created"`
	DepartmentId int            `json:"department_id" query:"department_id"`
	PositionId   int            `json:"posirion_id" query:"posirion_id"`
	CompanyId    int            `json:"company_id" query:"company_id"`
}

func NewApplication(db *sql.DB) *Handler {
	return &Handler{db}
}

type Err struct {
	Message string `json:"message"`
}

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "feelme_admin:zj439$Z2p@tcp(119.59.96.90)/feelme_db")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	return db
}
