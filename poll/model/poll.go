package model

type CreatePollRequest struct {
	Template bool `json:"template"`
}

type Poll struct {
	SessionId string     `json:"-"`
	Link      string     `json:"link"`
	Deadline  int        `json:"deadline,omitempty"`
	Questions []Question `json:"questions"`
}

type Link struct {
	Link string `json:"link" bson:"_id"`
}
