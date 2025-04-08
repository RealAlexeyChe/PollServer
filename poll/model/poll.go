package model

type CreatePollRequest struct {
	Course    string `json:"course"`
	Group     string `json:"group"`
	Professor string `json:"professor"`
	Deadline  int64  `json:"deadline,omitempty"`
}

type Poll struct {
	Link      string
	Course    string
	Group     string
	Professor string
	Deadline  int64 `json:"deadline,omitempty"`
	Questions []Question
}
