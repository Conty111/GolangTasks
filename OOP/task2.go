/*
Новый пользователь
Вам нужно создать структуру "Пользователь" (User) с следующими полями:

Имя (Name) Возраст (Age) Активный (Active) Поле "Имя" должно быть обязательным для заполнения при создании экземпляра структуры. Поле "Возраст" должно иметь значение по умолчанию 18, и поле "Активный" должно иметь значение по умолчанию true.

Требуется создать конструктор для структуры "Пользователь", который позволит инициализировать поля "Имя" и "Возраст", а поле "Активный" будет иметь значение по умолчанию.

Примечания
Код программы должен содержать описание струкрутры User:
type User struct { Name string Age int Active bool }
*/

package OOP

type User1 struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(name string, age int, active bool) *User1 {
	if name == "" {
		return nil
	}
	if age <= 0 {
		age = 18
	}
	return &User1{
		Name:   name,
		Age:    age,
		Active: true,
	}
}