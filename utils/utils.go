package utils

import (
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
)

func LogJsonRecieved(o any) {

	str, _ := json.Marshal(o)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	f := colorjson.NewFormatter()
	f.Indent = 2

	s, _ := f.Marshal(obj)
	fmt.Println("Получен JSON: ")
	fmt.Println(string(s))
}

func LogJsonLight(o any) {
	str, _ := json.Marshal(o)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	f := colorjson.NewFormatter()
	f.Indent = 2

	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}

func LogJsonSent(o any) {
	str, _ := json.Marshal(o)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	f := colorjson.NewFormatter()
	f.Indent = 2

	s, _ := f.Marshal(obj)
	fmt.Println("Отправлен JSON: ")
	fmt.Println(string(s))
}
