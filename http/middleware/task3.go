/*
Basic авторизация

Базовая HTTP-авторизация - это способ защитить веб-страницы или другие ресурсы от несанкционированного доступа.
Когда вы пытаетесь получить доступ к защищенному ресурсу, сервер запрашивает имя пользователя и пароль.
Если они корректные, то вы получаете доступ к ресурсу. Например, если у вас есть сайт
с персональными данными, то базовая HTTP-авторизация поможет защитить эти данные от злоумышленников.
При использовании базовой авторизации поверх HTTP используется заголовок Authorization

Напишите веб-сервер с использованием базовой HTTP-авторизации на пути /answer/.
Сервер должен проверять наличие и корректность заголовка Authorization и возвращать
ответ The answer is 42 при успешной авторизации.

При запросе без заголовка Authorization сервер должен вернуть статус 403 StatusForbidden и запросить авторизацию.

При запросе с заголовком Authorization, но некорректными данными для аутентификации сервер
должен вернуть статус 401 StatusUnauthorized и запросить авторизацию.

Middleware функцию назовите Authorization(http.HandlerFunc)

# Примечания

Корректные данные для пользователя
Login: userid
Password: password

Пример ответа:

POST /api/users HTTP/1.1 # метод и URL
Host: example.com # обязательный заголовок Host
Content-Type: application/json  # заголовок с типом данных
Authorization: Basic userid:password

	{
	  "name": "John Doe",
	  "email": "johndoe@example.com",
	  "password": "123456"
	} # тело запроса

Успешная авторизация:

curl -X GET http://127.0.0.1:5000/answer/ -H "Authorization: Basic dXNlcmlkOnBhc3N3b3Jk"
Ожидаемый ответ: The answer is 42

Отсутствие или некорректная авторизация: Пример запроса (с использованием cURL без заголовка Authorization):

curl -X GET http://127.0.0.1:5000/answer/
Ожидаемый ответ: ""
Ожидаемый статус HTTP: 403

Пример запроса (с использованием cURL с некорректными данными):

curl -X GET http://127.0.0.1:5000/answer/ -H "Authorization: Basic ZFhObDNtNWs1bjdoYzNOM2IzSms="
Ожидаемый ответ: ""
Ожидаемый статус HTTP: 401
*/

package middleware

import (
	"fmt"
	"net/http"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Host", "example.com")
		if auth == "Basic dXNlcmlkOnBhc3N3b3Jk" {
			next.ServeHTTP(w, r)
		} else if auth != "" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	})
}

func Answer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The answer is 42")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/answer/", Authorization(Answer))

	http.ListenAndServe("localhost:5000", mux)
}
