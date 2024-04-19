/*
Middleware Context

Вам необходимо создать веб-сервер с Middleware RequestIDMiddleware(next http.Handler) http.Handler
для HTTP-обработчика HelloHandler(w http.ResponseWriter, r *http.Request), который будет добавлять
информацию из заголовка "X-Request-ID" в контекст запроса и затем использовать эту информацию в самом обработчике.
Если "X-Request-ID" не передается - необходимо написать об этом в формате в виде "RequestID not found".

# Не забудьте про функцию StartContext, которая должна содержать методы для запуска сервера в виде

	func StartContext() {
	    http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))

	    http.ListenAndServe(":8080", nil)
	}

# Примечания

Запрос:

	HTTP метод: GET
	Путь: "/"
	Заголовок: X-Request-ID: 12345

Ожидаемый ответ:

	Статус: 200 OK
	Тело ответа: "Hello! RequestID: 12345"

Запрос:

	HTTP метод: GET
	Путь: "/"

Ожидаемый ответ:

	Статус: 200 OK
	Тело ответа: "Hello! RequestID not found"
*/
package context

import (
	"context"
	"fmt"
	"net/http"
)

const key = "X-Request-ID"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	val := r.Context().Value(key).(string)
	if val != "" {
		fmt.Fprintf(w, "Hello! RequestID: "+val)
	} else {
		fmt.Fprintf(w, "Hello! RequestID not found")
	}
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get(key)
		ctx := context.WithValue(r.Context(), key, val)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func StartContext() {
	http.Handle("/", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))

	http.ListenAndServe(":8080", nil)
}
