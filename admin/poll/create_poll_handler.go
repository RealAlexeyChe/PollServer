package poll

import (
	"LesyaBack/db"
	. "LesyaBack/poll/model"
	. "LesyaBack/utils"
	"encoding/json"
	"fmt"

	"net/http"
)

func HandleCreatePoll(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Запрос POST на /admin/poll")
	var r CreatePollRequest
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("Неправильный запрос")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	LogJsonRecieved(r)

	c, e := req.Cookie("sessionId")
	if e != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	l, e := CreatePoll(&r, c.Value)
	if e != nil {
		fmt.Println("Ошибка создания опроса: ", e)

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, e := json.Marshal(l)
	if e != nil {
		fmt.Println("Ошибка получения ссылки: ", e)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Отправлена ссылка на опрос: ", *l)

	rw.Write(res)

}

func CreatePoll(request *CreatePollRequest, sessionId string) (*Poll, error) {
	p, err := db.CreateNewPoll(request, sessionId)
	if err != nil {
		return nil, err
	}
	return p, nil
}
