package poll

import (
	"LesyaBack/db"
	. "LesyaBack/poll/model"
	. "LesyaBack/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandleUpdatePoll(rw http.ResponseWriter, req *http.Request) {

	link := strings.Replace(req.URL.Path, "/poll/", "", 1)
	link = link
	fmt.Println("Запрос PUT на /admin/poll/", link)

	var p Poll
	json.NewDecoder(req.Body).Decode(&p)

	LogJsonRecieved(p)
	res, err := db.UpdatePoll(link, p)
	if err != nil {
		fmt.Println("Опрос не найден")

		rw.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Внутренняя ошибка сервера: ", err)

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	LogJsonSent(res)
	rw.Write(b)
}
