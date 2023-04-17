package actions

import "database/sql"

const (
	createLogAttackDamage = "INSERT INTO logs (amount,datetime,account_id) VALUES ( ?, ?, ?) RETURNING log_id;"
)

type Handler struct {
	DB *sql.DB
}
type AttackDamageSender struct {
	Amount int `json:"amount" query:"amount"`
}

func NewApplication(db *sql.DB) *Handler {
	return &Handler{db}
}

type Err struct {
	Message string `json:"message"`
}
