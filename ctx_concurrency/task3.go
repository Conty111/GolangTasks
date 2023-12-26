package ctx_concurrency

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func fetch(url string, c chan *APIResponse, ctx context.Context) {
	var resp APIResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	cl := &http.Client{}
	r, err := cl.Do(req)
	if err != nil {
		resp.URL = url
		resp.Err = err
		c <- &resp
		return
	}
	defer r.Body.Close()
	var data []byte
	buf := make([]byte, 10)
	for {
		n, err := r.Body.Read(buf)
		data = append(data, buf[:n]...)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		if n < 10 {
			break
		}
	}

	resp.Data = string(data)
	resp.URL = url
	resp.StatusCode = r.StatusCode

	c <- &resp
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	res := make([]*APIResponse, len(urls))
	c := make(chan *APIResponse)
	for _, elem := range urls {
		go fetch(elem, c, newCtx)
	}
	for i := 0; i < len(urls); i++ {
		val := <-c
		res[i] = val
	}
	return res
}
