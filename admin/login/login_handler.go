package login

import (
	. "LesyaBack/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Запрос GET на /admin/login")
	var l LoginRequest
	err := json.NewDecoder(req.Body).Decode(&l)

	if err != nil {
		fmt.Println("Неправильный запрос")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	LogJsonRecieved(l)

	if l.Login == "lesya" && l.Password == "leti" {
		fmt.Println("Пользователь авторизован")
		rw.WriteHeader(http.StatusOK)
		return
	} else {
		fmt.Println("Неправильный логин или пароль")
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
}
