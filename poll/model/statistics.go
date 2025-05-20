package model

type Statistic struct {
	Text       string    `json:"text"`
	Votes      []float32 `json:"-"`
	Percentage []float64 `json:"percentage"`
}

func GetStats(p *[]Poll, a *[]Answer) []Statistic {
	total := len(*a)
	var res []Statistic = make([]Statistic, total)
	for _, s := range res {
		s.Votes := make([]float64, total)
	}

	for i, ans := range *a {
		if ans.Select != nil {
			var rec []string
			rec = append(rec, poll.Questions[i].Text)
			rec = append(rec, string(p))
		}

		var res Statistic
	}
}
