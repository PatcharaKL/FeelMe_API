package humanresource

import "database/sql"

const (
	getLogTimeStamp         = "SELECT log_timestamps.username,log_timestamp_types.name,log_timestamps.time FROM feelme_db.log_timestamps JOIN feelme_db.log_timestamp_types on feelme_db.log_timestamps.timestamp_type = feelme_db.log_timestamp_types.id WHERE  log_timestamps.user_id = ? ORDER BY log_timestamps.time ;"
	getUserFullNameByUserId = "SELECT name,surname FROM feelme_db.accounts WHERE account_id=? ;"
	editUserProfile         = "UPDATE feelme_db.accounts SET name = ? , surname = ? , position_id = ?, department_id = ? WHERE account_id = ?;"
	createdEmployee         = "INSERT INTO accounts (email, password_hash, name, surname, avatar_url, apply_date, is_active, hp, level, created, department_id, position_id, company_id) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) RETURNING account_id;"
)

type Handler struct {
	DB *sql.DB
}

func NewApplication(db *sql.DB) *Handler {
	return &Handler{db}
}

type Err struct {
	Message string `json:"message"`
}
