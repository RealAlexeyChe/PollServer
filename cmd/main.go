package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
import (
	. "LesyaBack/admin/login"
	. "LesyaBack/admin/poll"
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

func run() error {
	r := chi.NewRouter()

	r.Get("/info", infoHandler)
	r.Post("/admin/login", LoginHandler)

	r.Get("/admin/poll", HandleGetAllPolls)
	r.Post("/admin/poll", HandleCreatePoll)
	r.Put("/admin/poll", HandleUpdatePoll)
	r.Delete("/admin/poll/*", HandleDeletePoll)

	r.Get("/admin/poll/*", HandleGetPoll)
	fmt.Println("Сервер запущен на 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Не могу запустить сервер: ", err)
	}

	return nil
}
