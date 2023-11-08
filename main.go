package main

import (
	"github.com/Conty111/GolangTasks/sortirovki"
	"log"
)

func main() {
	//var x float32 = 1
	//log.Print(x / -0)
	//arr := []string{"Vasya", "Vasyaa", "Anatol", "Anna", "Bob"}
	//sortirovki.SortNames(arr)
	workers := []sortirovki.Worker{
		{Name: "Михаил", Position: "директор", Salary: 200, ExperienceYears: 5},
		{Name: "Игорь", Position: "зам. директора", Salary: 180, ExperienceYears: 3},
		{Name: "Николай", Position: "начальник цеха", Salary: 120, ExperienceYears: 2},
		{Name: "Андрей", Position: "мастер", Salary: 90, ExperienceYears: 10},
		{Name: "Виктор", Position: "рабочий", Salary: 80, ExperienceYears: 3},
	}
	zavod := sortirovki.Company{
		Workers: workers,
	}
	log.Print(zavod.SortWorkers())
}
