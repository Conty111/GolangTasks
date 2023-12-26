/*
Напишите функцию StartServer(maxTimeout time.duration) которая запусает веб-сервер
по адресу http://localhost:8080. При обращении по URL http://localhost:8080/readSource
сервер должен сделать запрос по другому адресу: http://localhost:8081/provideData
(код запуска сервера localhost:8081 писать не нужно) и вернуть полученные данные.
Используйте http.timeoutHandler для ограничения времени ожидания данных с сервера localhost:8081.
При привышении лимита maxTimeout пользователю должна веруться ошибка с кодом StatusServiceUnavailable, иначе - полученные данные.
*/

package ctx_concurrency

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type readSource struct{}

func (h *readSource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := http.NewRequest(r.Method, "localhost:8081/provideData", r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(buf)
}

func StartServer(maxTimeout time.Duration) {
	h := readSource{}
	var mux http.ServeMux
	mux.Handle("/readSource", http.TimeoutHandler(&h, maxTimeout, "StatusServiceUnavailable"))
	if err := http.ListenAndServe(":8080", &mux); err != nil {
		log.Fatal(err)
	}
}
