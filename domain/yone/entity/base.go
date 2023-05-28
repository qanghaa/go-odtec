package entity

import (
	"time"

	"github.com/jackc/pgtype"
)

type BaseEntity struct {
	CreatedAt pgtype.Timestamptz `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (e *BaseEntity) Now() {
	e.CreatedAt.Set(time.Now())
	e.UpdatedAt = e.CreatedAt
	e.DeletedAt.Set(nil)
}
