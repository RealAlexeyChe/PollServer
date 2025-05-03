package model

type Question struct {
	IsRequired bool
	Text       string       `json:"text"`
	Type       QuestionType `json:"type"`
	Rows       []string     `json:"grid,omitempty"`
	Options    []string     `json:"options,omitempty"`
}

type QuestionType string

const (
	Select      = "select"
	MultiSelect = "multiselect"
	Text        = "text"
	Grid        = "grid"
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

var TextExample Question = Question{
	Text: "Опиши свой день",
	Type: Text, // = "text"
}

var GridExample Question = Question{
	Text: "Оцени машины",
	Type: Grid, // = "gid"
	Options: []string{
		"Плохо",
		"Средне",
		"Хорошо",
	},
	Rows: []string{
		"Lexus RX",
		"Lada Grata",
		"Ford F-150",
	},
}
