/*
HTTP Клиент

Вам необходимо написать функцию SendHTTPRequest(url string) (string, error),
которая делает GET-запрос к заданному URL и возвращает тело ответа в виде строки.

Примечания
Нужный url подставляется внутри теста. В случае, если произошла ошибка,
необходимо возвращать ошибку: "Something went wrong..."

Пример ответа функции:

"Hello, World!\n"
*/

package context

import (
	"fmt"
	"io"
	"net/http"
)

func SendHTTPRequest(url string) (string, error) {
	myError := fmt.Errorf("Something went wrong...")
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, url, nil)
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
