CREATE TABLE IF NOT EXISTS exams (
  exam_id Text NOT NULL,
  name Text NOT NULL,
  grade_to_pass INTEGER NOT NULL DEFAULT 0,
  updated_at timestamptz NOT NULL,
	created_at timestamptz NOT NULL,
	deleted_at timestamptz,
  CONSTRAINT exams_pk PRIMARY KEY (exam_id)
)
 
CREATE TABLE IF NOT EXISTS public.quizzes (
	quiz_id text NOT NULL,
	question text NOT NULL,
  answers _text NOT NULL,
	correct_answer text NOT NULL,
  exam_ids _text NOT NULL,
	updated_at timestamptz NOT NULL,
	created_at timestamptz NOT NULL,
	deleted_at timestamptz,
	CONSTRAINT quizzes_pk PRIMARY KEY (quiz_id)
);

CREATE TABLE IF NOT EXISTS public.exam_submissions (
	exam_submission_id text NOT NULL,
  score Integer,
  student_id text NOT NULL,
  exam_id text NOT NULL,
  is_passed Boolean,
	updated_at timestamptz NOT NULL,
	created_at timestamptz NOT NULL,
	deleted_at timestamptz,
	CONSTRAINT exam_submissions_pk PRIMARY KEY (exam_submission_id)
);


