package entity

import (
	"go-odtec/utils/database"

	"github.com/jackc/pgtype"
)

type Quiz struct {
	QuizID        pgtype.Text      `json:"quiz_id,omitempty" db:"quiz_id"`
	Question      pgtype.Text      `json:"question,omitempty" db:"question"`
	Answers       pgtype.TextArray `json:"answers,omitempty" db:"answers"`
	CorrectAnswer pgtype.Text      `json:"correct_answer,omitempty" db:"correct_answer"`
	ExamIDs       pgtype.TextArray `json:"exam_id,omitempty" db:"exam_ids"`
	BaseEntity
}

func (q *Quiz) FieldMap() ([]string, []interface{}) {
	return database.FieldMap(q)
}

func (q *Quiz) TableName() string {
	return "quizzes"
}
