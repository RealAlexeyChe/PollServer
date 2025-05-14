package model

type AnswersForm struct {
	Link    string   `json:"link"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Select      *int    `json:"select,omitempty"`
	Multiselect *[]int  `json:"multiselect,omitempty"`
	Grid        *[]int  `json:"grid,omitempty"`
	Text        *string `json:"text,omitempty"`
}
