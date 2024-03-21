/*
Привет, странник!

Напишите веб-сервер, который при обращении к нему будет возвращать приветствие
с именем пользователя, полученным из параметра запроса.

Если параметр пустой или отсутствует, сервер должен вернуть приветствие hello stranger.
Если ответ содержит не только английские буквы, приветствие должно быть hello dirty hacker.

Веб-сервер должен быть запущен на порту с номером 8080.

Пример запроса:

curl localhost:8080/?name=John
# hello John

curl localhost:8080
# hello stranger

*/

package http_server

import (
	"fmt"
	"net/http"
)

func StartStranger() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) > 0 && check(name) {
			fmt.Fprint(w, "hello "+name)
		} else if len(name) > 0 {
			fmt.Fprint(w, "hello dirty hacker")
		} else {
			fmt.Fprint(w, "hello stranger")
		}
	})
	http.ListenAndServe(":8080", nil)
}

func check(txt string) bool {
	for _, letter := range txt {
		if !(letter <= 'Z' && letter >= 'A') && !(letter <= 'z' && letter >= 'a') {
			return false
		}
	}
	return true
}
