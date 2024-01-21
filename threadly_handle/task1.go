/*
Напишите функцию StringsGen(lines ...string) <-chan string,
которая принимает на вход набор строк и возвращает их через канал.
*/

package threadly_handle

func StringsGen(lines ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, l := range lines {
			out <- l
		}
	}()
	return out
}
