/*
Напишите функцию func FetchURL(url string) string Которая пойдет по адресу url.
В случае успеха функция должна вернуть "Successfully fetched"
В случае неудачи функция должна вернуть "Failed to fetch"
*/

package timeouts

import (
	"net/http"
)

func FetchURL(url string) string {
	cl := http.Client{}
	//c := make(chan string)
	response, err := cl.Get(url)
	if err != nil {
		return "Failed to fetch"
	}
	if response.StatusCode != 200 && response.StatusCode != 201 {
		return "Failed to fetch"
	}
	return "Successfully fetched"
}
