package poll

import (
	"LesyaBack/db"
	"fmt"
	"net/http"
	"strings"
)

func HandleDeletePoll(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Запрос DELETE на /admin/poll")

	link := strings.Replace(req.URL.Path, "/admin/poll/", "", 1)
	link = link

	c, e := req.Cookie("sessionId")
	if e != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := db.DeletePoll(link, c.Value)
	if err != nil {
		if err.Error() == "Not found" {
			fmt.Println("Опрос по ссылке не найден")
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println("Внутренняя ошибка сервера: ", err)

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Опрос ", link, "удалён")

	rw.WriteHeader(http.StatusNoContent)
}
