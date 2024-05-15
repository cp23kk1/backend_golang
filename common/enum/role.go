package enum

import "database/sql/driver"

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
	GUEST Role = "guest"
)

func (s *Role) Scan(value interface{}) error {
	*s = Role(value.([]byte))
	return nil
}

func (s Role) Value() (driver.Value, error) {
	return string(s), nil
}

type QuestionType string

const (
	VOCABULARY QuestionType = "vocabulary"
	SENTENCE   QuestionType = "sentence"
	PASSAGE    QuestionType = "passage"
)

func (s *QuestionType) Scan(value interface{}) error {
	*s = QuestionType(value.([]byte))
	return nil
}

func (s QuestionType) Value() (driver.Value, error) {
	return string(s), nil
}
