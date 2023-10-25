/*
Ошибочный странник

Давайте модифицируем самый первый сервер, который мы сделали. Да-да, тот самый, что отвечает hello stranger.

Переделайте обработку некорректного значения name в middleware Sanitize(http.HandlerFunc)
и обработку пустого значения в middleware SetDefaultName(http.HandlerFunc).

| Sanitize Middleware: Middleware "Sanitize" должен проверить значение параметра "name" и убедиться, что оно не содержит символов, которые не подходят для имени. Каждый запрос с некорректными символами в "name" должен возвращать статус 400 (Bad Request).
| SetDefaultName Middleware: Middleware "SetDefaultName" должен проверить, если параметр "name" пустой, и, если это так, установить значение "name" равным "stranger".

Вы должны реализовать оба middleware и добавить их к серверу, чтобы обеспечить корректное поведение.
Каждый middleware должен быть применен к URL-пути "/". Порт должен быть 8080

# Примечания

Обработка некорректного значения означает, что name не должно содержать символов, которые не подходят для имени.
То есть, оно не должно содержать символов !@#$%^&*()_+.

При получении в поле name пустой строки метод SetDefaultName должен устанавливать поле name равным "stranger".

Примеры запросов:

	Запрос "/?name=John" с корректным именем "John" должен вернуть "Hello, John!".
	Запрос "/?name=John!" с некорректными символами в имени должен вернуть статус 400 (Bad Request).
	Запрос "/?name=" с пустым именем должен вернуть "Hello, stranger!".

Ваша задача - реализовать и протестировать middleware "Sanitize" и "SetDefaultName" для сервера и удостовериться,
что они работают корректно с различными значениями параметра "name".
*/

package middleware

import (
	"net/http"
)

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) > 0 {
			if check(name) {
				w.Write([]byte(name))
			} else {
				w.Write([]byte("dirty hacker"))
			}
		}
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, "))

		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			w.Write([]byte("stranger"))
		}
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("!"))
}

func check(txt string) bool {
	for _, letter := range txt {
		if !(letter <= 'Z' && letter >= 'A') && !(letter <= 'z' && letter >= 'a') {
			return false
		}
	}
	return true
}
