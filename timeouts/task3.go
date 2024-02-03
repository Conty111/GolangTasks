/*


Напишите программу, имитирующую портал викторин,
там будет несколько предустановленных вопросов и ответов на них.
Вопросы будут приходить один за другим и будут ждать 15 секунд,
пользователь должен ввести ответ с клавиатуры в течение этого времени.
Если время истекло, отобразите сообщение о том, что время истекло, и переходите к следующему вопросу.
В конце покажите общий правильный ответ завершенного теста.
Сигнатура проверяемой функции func QuizRunner(questions, answers []string) int

questions - список вопросов answers - список ответов результат - число правильных ответов

*/

package timeouts

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func QuizRunner(questions, answers []string) int {
	score := 0
	ch := make(chan string)
	sc := bufio.NewScanner(os.Stdin)
	for idx := 0; idx < len(questions); idx++ {
		go func() {
			sc.Scan()
			ch <- sc.Text()
		}()
		select {
		case res := <-ch:
			if res == answers[idx] {
				score++
			} else {
				fmt.Println("Wrong answer!", res)
			}
		case <-time.After(time.Second * 5):
			fmt.Println("Time's out!")
		}
	}
	return score
}
