package login

import (
	"LesyaBack/db"
	"LesyaBack/poll/model"
	. "LesyaBack/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
)

func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Запрос GET на /admin/login")
	var a model.AdminRequest

	err := json.NewDecoder(req.Body).Decode(&a)
	LogJsonRecieved(req.Body)

	if err != nil {
		fmt.Println("Неправильный запрос")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = mail.ParseAddress(a.Email)
	if err != nil {
		fmt.Println("Неправильный запрос")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Неправильный адрес почты"))
		return
	}

	id := db.GetAdminId(&a)

	if id == nil {

		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Логин или пароль неправильные"))
		return
	}

	cookie := http.Cookie{
		Name:     "sessionId",
		Value:    *id,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(rw, &cookie)

}
