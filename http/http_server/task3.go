/*
Сервер Фибоначчи с метриками

Напишите веб-сервер, который будет считать метрики времени ответа сервиса.

Возьмите в качестве основы веб-сервер из предыдущей задачи, вычисляющий числа Фибоначчи,
и добавьте к нему хендлер /metrics, который отдаёт:

rpc_duration_milliseconds_count 10

где 10 — число раз, которое вызвали хендлер, отвечающий числами Фибоначчи. Выглядит это так:

curl http://localhost:8080/metrics
# rpc_duration_milliseconds_count 0

curl http://localhost:8080/
# 0

curl http://localhost:8080/metrics

# rpc_duration_milliseconds_count 1

curl http://localhost:8080/
# 34

curl http://localhost:8080/metrics
# rpc_duration_milliseconds_count 10

*/

package http_server

import (
	"fmt"
	"net/http"
)

var fib_metr1, fib_metr2 int = 0, 1
var i int = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fib1)
		a := fib_metr1 + fib_metr2
		fib_metr1 = fib_metr2
		fib_metr2 = a
		i += 1
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", i)
	})
	http.ListenAndServe(":8080", nil)
}
