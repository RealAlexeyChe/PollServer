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

	link := strings.Replace(req.URL.Path, "/admin/poll/", "", 1)
	link = link
	fmt.Println("Запрос PUT на /admin/poll/", link)

	var p Poll
	err := json.NewDecoder(req.Body).Decode(&p)

	if err != nil {
		fmt.Println("Неправильный запрос")
		rw.WriteHeader(http.StatusBadRequest)

	}

	LogJsonRecieved(p)

	c, e := req.Cookie("sessionId")
	if e != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	res, err := db.UpdatePoll(link, p, c.Value)

	if err != nil {
		if err.Error() == "Unauthorized" {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println("Не удалось обновить запрос: ", err)

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
