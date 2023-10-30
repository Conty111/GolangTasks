/*
Управление транспортом
Вам нужно создать систему управления транспортными средствами, такими как автомобили и мотоциклы.
Каждое транспортное средство должно иметь метод для расчета времени в пути до определенного пункта назначения.

Создайте интерфейс Vehicle, который будет представлять транспортное средство и иметь метод
CalculateTravelTime(distance float64) float64 для расчета времени в пути.

Реализуйте две структуры, Car (автомобиль) и Motorcycle (мотоцикл), обе должны реализовывать
интерфейс Vehicle и иметь соответствующие поля для хранения данных о транспортных средствах
(например, скорость и тип транспортного средства).

Создайте функцию EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64, которая принимает список
транспортных средств и расстояние до пункта назначения, а затем возвращает карту (map),
где ключи - это типы транспортных средств, а значения - время в пути для каждого транспортного средства.
Используйте полиморфизм для вызова метода CalculateTravelTime() на каждом транспортном средстве независимо от его типа.
*/

package OOP

import "fmt"

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Speed    float64
	Type     string
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Speed    float64
	Type     string
	FuelType string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	res := make(map[string]float64, len(vehicles))
	for _, elem := range vehicles {
		res[fmt.Sprintf("%T", elem)] = elem.CalculateTravelTime(distance)
	}
	return res
}
