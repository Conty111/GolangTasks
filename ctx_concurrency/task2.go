/*
Напишите функцию fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error),
которая запрашивает данные по адресу url (метод GET) и возвращает код ответа и само тело ответа.
Используйте контекст для ограничения времени запроса и отмены ожидания свыше timeout.
В случае возникновения ошибок - возвращайте nil, error. При превышении таймаута ожидания - верните nil, context.DeadlineExceeded.

Примечания
Код должет содержать структуру:
type APIResponse struct {
	Data string // тело ответа
	StatusCode int // код ответа
}
*/

package ctx_concurrency

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type APIResponse1 struct {
	Data       string
	StatusCode int
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse1, error) {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(newCtx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	c := &http.Client{}
	r, err := c.Do(req)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		if n < 10 {
			break
		}
	}
	ans := APIResponse1{
		StatusCode: r.StatusCode,
		Data:       string(data),
	}
	return &ans, nil
}
