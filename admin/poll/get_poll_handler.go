package poll

import (
	"LesyaBack/db"
	"LesyaBack/poll/model"
	. "LesyaBack/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandleGetAllPolls(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Запрос GET на /admin/poll")

	var p []model.Poll
	p, err := db.GetAllPolls()
	if err != nil {
		fmt.Println("Внутренняя ошибка сервера: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Внутренняя ошибка сервера: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Список опросов:")
	for _, l := range p {
		LogJsonLight(l)
	}
	rw.Write(res)
}

func HandleGetPoll(rw http.ResponseWriter, req *http.Request) {

	link := strings.Replace(req.URL.Path, "/admin/poll/", "", 1)

	fmt.Println("Запрос GET на /admin/poll/", link)
	res, err := db.GetPoll(link)
	if err != nil {
		fmt.Println("Опрос не найден")
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	p, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Внутренняя ошибка сервера: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	LogJsonSent(*res)
	rw.Write(p)
}
