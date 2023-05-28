package entity

import (
	"go-odtec/utils/database"

	"github.com/jackc/pgtype"
)

type ExamSubmission struct {
	ExamSubmissionID pgtype.Text `json:"exam_submission_id,omitempty" db:"exam_submission_id"`
	Score            pgtype.Text `json:"score,omitempty" db:"score"`
	StudentID        pgtype.Text `json:"student_id,omitempty" db:"student_id"`
	ExamID           pgtype.Text `json:"exam_id,omitempty" db:"exam_id"`
	IsPassed         pgtype.Bool `json:"is_passed,omitempty" db:"is_passed"`
	BaseEntity
}

func (s *ExamSubmission) FieldMap() ([]string, []interface{}) {
	return database.FieldMap(s)
}

func (s *ExamSubmission) TableName() string {
	return "exam_submissions"
}
