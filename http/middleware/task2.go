/*
Обработка метрик

Продолжаем развивать наши skill'ы!

Возьмите сервер Фибоначчи с метриками из предыдущего урока.
Переделайте увеличение метрики числа запросов в middleware Metrics(http.HandlerFunc) и
добавьте этот middleware для url /.Metrics будет измерять время выполнения каждого запроса
и увеличивать счетчик запросов. Этот middleware должен быть добавлен к обработке всех запросов, включая /.
Создайте отдельный обработчик /metrics, который будет возвращать метрику времени выполнения
запросов в формате строки, например: rpc_duration_milliseconds_count 10.
Это число должно увеличиваться с каждым новым запросом.

# Примечания

Функцию main создавать не надо.
Нумерация элементов в программировании всегда начинается с 0.

Для запросов к / (Фибоначчи):

	Ответ с числом Фибоначчи, например:
	    Запрос: GET /
	    Ответ: "5".

Для запросов к /metrics (Метрики):

	    Ответ с метриками сервера, включая количество запросов
		(число запросов, увеличивающееся с каждым запросом), например:
	        Запрос: GET /metrics
	        Ответ:
	        rpc_duration_milliseconds_count 5
	        Где "5" - это количество запросов, которые были обработаны сервером.

Важно: Ваша программа должна содержать следующий код (используйте как шаблон для своей программы):

package main

import (

	"net/http"

)
var requestCount int

	func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
		requestCount++
	}

func MetricsHandler(w http.ResponseWriter, r *http.Request) { }
*/
package middleware

import (
	"fmt"
	"net/http"
)

var f1, f2 int = 0, 1
var requestCount int = 0

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	fmt.Fprint(w, f1)
	a := f1 + f2
	f1 = f2
	f2 = a
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}
