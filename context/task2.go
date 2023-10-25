/*
HTTP Клиент с контекстом

Используя http.Client и context, напишите функцию
SendHTTPRequestWithContext(ctx context.Context, url string) (string, error),
которая делает GET-запрос к заданному URL и принимает контекст для управления запросом.

Данная задача похожа на предыдущую, однако следует добавить контекст в функцию,
которая выполняет запрос. Ошибки обрабатывайте аналогично.
Для выполнения запроса используйте функцию NewRequestWithContext.

Пример возвращаемой функцией строки:
"Hello, World!\n"
*/

package context

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func SendHTTPRequestWithContext(ctx context.Context, url string) (string, error) {
	myError := fmt.Errorf("Something went wrong...")
	client := &http.Client{}
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := client.Do(r)
	if err != nil {
		return "", myError
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", myError
	}
	return string(respBody), nil
}
