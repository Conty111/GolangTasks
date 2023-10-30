/*
Отчетики
Вам нужно создать программу для создания отчетов о пользователях. У вас есть структура User
с данными о пользователях и структура Report, которая встраивает в себя структуру User.
Ваша задача - создать отчеты о пользователях на основе их данных.

Создайте структуру User со следующими полями:

ID (уникальный идентификатор пользователя)
Name (имя пользователя)
Email (электронная почта пользователя)
Age (возраст пользователя)

Создайте структуру Report, которая встраивает в себя структуру User и добавляет следующие поля:

ReportID (уникальный идентификатор отчета)
Date (дата создания отчета)

Создайте функцию CreateReport(user User, reportDate string) Report,
которая принимает пользователя и дату и возвращает отчет с заполненными данными.
Уникальный ReportID можно сгенерировать, например, с использованием времени.

Создайте функцию PrintReport(report Report), которая выводит информацию из отчета,
включая данные о пользователе и дату создания отчета.

Создайте функцию GenerateUserReports(users []User, reportDate string) []Report,
которая принимает список пользователей и дату и возвращает список отчетов о пользователях.
Для каждого пользователя в списке создайте отчет, используя функцию CreateReport,
и добавьте его в результирующий список.
*/

package OOP

import (
	"fmt"
	"math"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	return Report{
		ReportID: int(math.Pi*float64(time.Now().UnixNano())) % 1000000,
		Date:     reportDate,
		User:     user,
	}
}

func PrintReport(report Report) {
	fmt.Printf("%s\n%d\n", report.Name, report.ID)
}

func GenerateUserReports(users []User, reportDate string) []Report {
	res := make([]Report, len(users))
	for idx, user := range users {
		res[idx] = CreateReport(user, reportDate)
	}
	return res
}
