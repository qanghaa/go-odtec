package entity

import (
	"go-odtec/utils/database"

	"github.com/jackc/pgtype"
)

type Exam struct {
	ExamID      pgtype.Text `json:"exam_id,omitempty" db:"exam_id"`
	Name        pgtype.Text `json:"name,omitempty" db:"name"`
	GradeToPass pgtype.Int2 `json:"grade_to_pass,omitempty" db:"grade_to_pass"`
	BaseEntity
}

func (e *Exam) FieldMap() ([]string, []interface{}) {
	return database.FieldMap(e)
}

func (e *Exam) TableName() string {
	return "exams"
}
