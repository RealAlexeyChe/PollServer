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

func CreateAdminHandler(rw http.ResponseWriter, req *http.Request) {
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
		fmt.Println("Неправильный адрес почты")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Неправильный адрес почты"))
		return
	}

	if db.GetAdminId(&a) != nil {
		rw.Write([]byte("Админ уже существует!"))
		return
	}

	err = db.CreateAdmin(&a)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.Write([]byte("Админ создан!"))
}
