package model

type Question struct {
	Text     string       `json:"text"`
	Type     QuestionType `json:"type"`
	Options  string       `json:"options,omitempty"`
	MaxGrade int          `json:"maxGrade,omitempty"`
}

type QuestionType string

const (
	Select      = "select"
	MultiSelect = "multiselect"
	Grade       = "grade"
	Text        = "text"
)
