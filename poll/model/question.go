package model

type Question struct {
	Text     string       `json:"text"`
	Type     QuestionType `json:"type"`
	Options  []string     `json:"options,omitempty"`
	MaxGrade int          `json:"maxGrade,omitempty"`
}

type QuestionType string

const (
	Select      = "select"
	MultiSelect = "multiselect"
	Grade       = "grade"
	Text        = "text"
)

var SelectExample Question = Question{
	Text: "Выбери одну из опций",
	Type: Select, // = "select"
	Options: []string{
		"мало",
		"средне",
		"много",
	},
}

var MultiSelectExample Question = Question{
	Text: "Выбери одну или нескольно блюд",
	Type: MultiSelect, // = "multiselect"
	Options: []string{
		"пицца",
		"паста",
		"тирамису",
	},
}

var GradeExample Question = Question{
	Text:     "Дай оценку",
	Type:     Grade, // = "grade"
	MaxGrade: 10,
}

var TextExample Question = Question{
	Text: "Опиши свой день",
	Type: Text, // = "text"
}
