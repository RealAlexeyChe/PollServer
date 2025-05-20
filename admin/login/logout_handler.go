package login

import (
	"fmt"
	"net/http"
)

func LogoutHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Запрос GET на /admin/login")

	_, err := req.Cookie("sessionId")
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(rw, "Куки не установлено, выйти невозможно", http.StatusUnauthorized)
	}
	cookie := http.Cookie{
		Name:     "sessionId",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(rw, &cookie)
	rw.Write([]byte("Выход выполнен, куки удалена"))
	rw.WriteHeader(http.StatusOK)

}
