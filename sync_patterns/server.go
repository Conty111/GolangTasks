package sync_patterns

import (
	"log"
	"net/http"
)

func Start() {
	var mux http.ServeMux
	mux.HandleFunc("/mark", func(w http.ResponseWriter, r *http.Request) {
		students := make(map[string]byte)
		students["test"] = 5
		students["bad"] = 2
		students["good"] = 4
		name := r.URL.Query().Get("name")
		res, ok := students[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte{res})
	})
	err := http.ListenAndServe(":8082", &mux)
	if err != nil {
		log.Fatal(err)
	}
}
