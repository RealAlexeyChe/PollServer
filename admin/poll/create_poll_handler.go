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
	if r.Course == "" || r.Group == "" || r.Professor == "" {
		fmt.Println("Не заполнена одно из полей")
	}
	if r.Course == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Course must be filled"))
		return
	}
	if r.Group == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Group must be filled"))
		return
	}
	if r.Professor == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Professor must be filled"))
		return
	}
	l, e := CreatePoll(&r)
	if e != nil {
		fmt.Println("Ошибка создания опроса: ", e)

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, e := json.Marshal(l)
	fmt.Println("Отправлена ссылка на опрос: ", *l)
	if e != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write(res)

}

func CreatePoll(request *CreatePollRequest) (*string, error) {
	l, err := db.CreateNewPoll(request)
	if err != nil {
		return nil, err
	}
	return l, nil
}
