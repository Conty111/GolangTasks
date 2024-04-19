/*
Сервер Фибоначчи

Напишите веб-сервер, который будет запускаться на 8080 порту и принимать запросы
на получение следующего числа Фибоначчи, возвращая его значение.

Примеры запросов и ответов:

curl http://localhost:8080/
# 0

curl http://localhost:8080/
# 1

curl http://localhost:8080/
# 1

curl http://localhost:8080/
# 2

Сервер не сохраняет свое состояние между перезапусками.
Таким образом, если закрыть программу и запустить ее заново - подсчет начнется с 0.

*/

package http_server

import (
	"fmt"
	"net/http"
)

var fib1, fib2 int = 0, 1

func StartFib() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fib1)
		a := fib1 + fib2
		fib1 = fib2
		fib2 = a
	})
	http.ListenAndServe(":8080", nil)
}
