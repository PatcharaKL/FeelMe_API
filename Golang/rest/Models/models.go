package models

import (
	"database/sql"
)

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