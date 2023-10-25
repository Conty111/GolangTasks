/*
Обогащение ответа страннику
Возьмите сервер, отвечающий hello stranger из этого урока. Переделайте формат ответа на JSON вида:

{"greetings": "hello", "name": "stranger"}

При этом name берётся из запроса, а логика его подстановки не меняется.

Добавьте к этому middleware RPC(http.HandlerFunc), которая заменяет ответ на формат:

{"status": "ok", "result": {"greetings": "hello", "name": "stranger"}}

Так же, переделайте Middleware Sanitize, чтобы она возвращала panic в случае некорректного имени,
и добавьте обработку этой паники в новой middleware RPC так, чтобы отдавать пользователю ответ:

{"status": "error", "result": {}}

# Примечания

Запрос: GET /hello?name=Alice
Ответ: {"greetings": "Hello","name": "Alice"}

Запрос: GET /hello?name=""

Ответ: Ошибка 500, так как имя не соответствует требованиям middleware Sanitize.
Ответ: {"status": "error","result": {}}
*/
package middleware

import (
	"fmt"
	"net/http"
)

func Sanitize1(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if !check1(name) {
			panic(fmt.Errorf("dirty hacker"))
		}
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName1(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			value := r.URL.Query()
			value.Set("name", "stranger")
			r.URL.RawQuery = value.Encode()
		}
		next.ServeHTTP(w, r)
	})
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Write([]byte(`{"status":"error","result":{}}`))
				return
			}
			name := r.URL.Query().Get("name")
			res := fmt.Sprintf(`{"greetings":"hello","name":"%s"}`, name)
			w.Write([]byte(fmt.Sprintf(`{"status":"ok","result":%s}`, res)))
		}()
		next.ServeHTTP(w, r)
	})
}

func HelloHandler1(w http.ResponseWriter, r *http.Request) {}

func check1(txt string) bool {
	for _, letter := range txt {
		if !(letter <= 'Z' && letter >= 'A') && !(letter <= 'z' && letter >= 'a') {
			return false
		}
	}
	return true
}
