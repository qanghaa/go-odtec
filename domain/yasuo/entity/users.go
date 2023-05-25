package entity

import (
	"go-odtec/golibs/database"

	"github.com/jackc/pgtype"
)

type Address struct {
	Street  pgtype.Text `json:"street,omitempty" db:"street"`
	City    pgtype.Text `json:"city,omitempty" db:"city"`
	State   pgtype.Text `json:"state,omitempty" db:"state"`
	Country pgtype.Text `json:"country,omitempty" db:"country"`
}

type User struct {
	UserID   pgtype.Text        `json:"user_id,omitempty" db:"user_id"`
	Email    pgtype.Text        `json:"email,omitempty" db:"email"`
	Name     pgtype.Text        `json:"name,omitempty" db:"name"`
	Gender   pgtype.Text        `json:"gender,omitempty" db:"gender"`
	Avatar   pgtype.Text        `json:"avatar,omitempty" db:"avatar"`
	Birthday pgtype.Timestamptz `json:"birthday,omitempty" db:"birthday"`
	Address
	BaseEntity
}

func (u *User) FieldMap() ([]string, []interface{}) {
	return database.FieldMap(u)
}

func (u *User) TableName() string {
	return "users"
}
