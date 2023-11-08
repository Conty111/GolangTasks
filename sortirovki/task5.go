/*
Сортировка сотрудников предприятия

На предприятии работают несколько cсотрудников. Каждый из них имеет свою должность,
фиксированную месячную заработную плату, и стаж работы.
Напишите программу в котором тип Company реализует следующий интерфейс:

type CompanyInterface interface{
AddWokerInfo(name, position string, salary, experience uint) error
SortWorkers() ([]string, error)
}

AddWokerInfo - метод добавления информации о новых сотрудниках, где
name - имя сотрудника,
position - должность,
salary - месячная зароботная плата,
experience - стаж работы (месяцев).

SortWorkers - метод сортировки, который сортирует список сотрудников по следующим свойствам:
доход за время работы на предприятии(по убыванию),
должность (от высокой)
и возвращает слайсл формата: имя - доход - должность.
Допустимые должности в порядке убывания: "директор", "зам. директора", "начальник цеха", "мастер", "рабочий".

Примечания

Пример отсортированного вывода:
Михаил - 12000 - директор
Андрей - 11800 - мастер
Игорь - 11000 - зам. директора
*/

package sortirovki

import (
	"fmt"
	"slices"
)

var positions = []string{"директор", "зам. директора", "начальник цеха", "мастер", "рабочий"}

type CompanyInterface interface {
	AddWokerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct {
	Name            string
	Position        string
	Salary          uint
	ExperienceYears uint
}

type Company struct {
	Workers []Worker
}

func (z *Company) AddWokerInfo(name, position string, salary, experience uint) error {
	for _, pos := range positions {
		if pos == position {
			z.Workers = append(z.Workers, Worker{
				Name:            name,
				Position:        position,
				Salary:          salary,
				ExperienceYears: experience,
			})
		}
	}
	return fmt.Errorf("Positions is not available")
}

func (z *Company) SortWorkers() ([]string, error) {
	slices.SortFunc(z.Workers, func(a, b Worker) int {
		var res int
		res = checkSalary(&a, &b)
		if res == 0 {
			return checkPosition(&a, &b)
		}
		return res
	})
	res := make([]string, len(z.Workers))
	for idx, elem := range z.Workers {
		row := fmt.Sprintf("%s - %d - %s", elem.Name, elem.Salary*elem.ExperienceYears*12, elem.Position)
		res[idx] = row
	}
	return res, nil
}

func checkSalary(a, b *Worker) int {
	if a.Salary*(a.ExperienceYears) > b.Salary*(b.ExperienceYears) {
		return -1
	} else if a.Salary*(a.ExperienceYears) < b.Salary*(b.ExperienceYears) {
		return 1
	}
	return 0
}

func checkPosition(a, b *Worker) int {
	var i, j int = -1, -1
	for idx, pos := range positions {
		if i != -1 && j != -1 {
			break
		}
		if pos == a.Position {
			i = idx
		}
		if pos == b.Position {
			j = idx
		}
	}
	if i > j {
		return 1
	} else if i < j {
		return -1
	}
	return 0
}

//[Михаил - 12000 - директор Андрей - 10800 - мастер Игорь - 6480 - зам. директора Николай - 2880 - начальник цеха Виктор - 2880 - рабочий]
//[Михаил - 12000 - директор Андрей - 10800 - мастер Николай - 2880 - начальник цеха Виктор - 2880 - рабочий]
//[Михаил - 12000 - директор Андрей - 10800 - мастер Николай - 2880 - начальник цеха Виктор - 2880 - рабочий]
