package poll

import (
	"LesyaBack/db"
	. "LesyaBack/poll/model"
	. "LesyaBack/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleCreateAnswers(rw http.ResponseWriter, req *http.Request) {
	var r AnswersForm
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("Неправильный запрос")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	LogJsonRecieved(r)
	p, err := db.GetPoll(r.Link)
	if err != nil {
		fmt.Println("Ошибка загрузки опроса по ссылке")
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	if len(r.Answers) != len(p.Questions) {
		fmt.Println("Количество вопросов и ответов не совпадает:")
		fmt.Println("Вопростов: ", len(p.Questions))
		fmt.Println("Ответов: ", len(r.Answers))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, a := range r.Answers {
		q := p.Questions[i]
		if q.Type == Select {
			if a.Select == nil {
				ReportAnswerTypeMistach(a, q)
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		if q.Type == MultiSelect {
			if a.Multiselect == nil {
				ReportAnswerTypeMistach(a, q)
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		if q.Type == Grid {
			if a.Grid == nil {
				ReportAnswerTypeMistach(a, q)
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		if q.Type == Text {
			if a.Text == nil {
				ReportAnswerTypeMistach(a, q)
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
	fmt.Println("Проверка типов ответов пройдена")

}
