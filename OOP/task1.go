/*
Книжная полка
Вы должны создать структуру "Книга" (Book) с следующими полями:

Название (Title) Автор (Author) Год выпуска (Year) Жанр (Genre) Требуется создать конструктор для структуры "Книга", который позволит инициализировать поля структуры при создании экземпляра. Конструктор должен принимать значения для всех полей и возвращать указатель на созданный экземпляр структуры "Книга".

Примечания
Код программы должен содержать описание струкрутры Book:
type Book struct { Title string Author string Year int Genre string }
*/

package OOP

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func NewBook(title string, author string, year int, genre string) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Year:   year,
		Genre:  genre,
	}
}
