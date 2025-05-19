package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
import (
	. "LesyaBack/admin/login"
	. "LesyaBack/admin/poll"
	"LesyaBack/poll/model"
	. "LesyaBack/user"
	. "LesyaBack/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	if e := run(); e != nil {
		log.Fatal(e)
	}
}

func infoHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Запрос GET на /info")
	rw.Write([]byte("LETI Polls API"))

}
func pathHandler(rw http.ResponseWriter, req *http.Request) {
	p := req.URL
	m := req.Method
	fmt.Println("Запрос " + m + " на " + p.Path)
	LogJsonRecieved(req.Body)
	fmt.Println("Ручка не найдена")

	http.Error(rw, "Путь не существует", http.StatusNotFound)
}
func run() error {

	r := chi.NewRouter()
	// ОТВЕТЫ
	r.HandleFunc("/*", pathHandler)
	r.Post("/answers/create", HandleCreateAnswers)
	// ОПРОСЫ

	r.Get("/info", infoHandler)
	r.Post("/admin/create", CreateAdminHandler)
	r.Post("/admin/login", LoginHandler)

	r.Get("/admin/poll", HandleGetAllPolls)
	r.Post("/admin/poll/create", HandleCreatePoll)
	r.Put("/admin/poll/*", HandleUpdatePoll)
	r.Delete("/admin/poll/*", HandleDeletePoll)

	r.Get("/admin/poll/*", HandleGetPoll)
	fmt.Println("Сервер запущен на 3000")

	fmt.Println("Образцы вопросов:")

	fmt.Println("Выбор одного из ответов:")
	LogJsonLight(model.SelectExample)
	fmt.Println("Выбор нескольких ответов:")

	LogJsonLight(model.MultiSelectExample)

	fmt.Println("Текстовая форма:")

	LogJsonLight(model.TextExample)

	fmt.Println("Сетка вопросов:")

	LogJsonLight(model.GridExample)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Не могу запустить сервер: ", err)
	}

	return nil
}
