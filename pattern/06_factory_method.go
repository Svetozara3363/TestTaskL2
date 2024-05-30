package main

import "fmt"

/*Фабричный метод — это порождающий паттерн проектирования, который определяет
общий интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

Плюсы:
- Избавляет от привязки к конкретному объекту
- Отделяет производство объектов

Минусы:
- Может привести к созданию больших параллельных иерархий объектов
- Один "божественный" конструктор
- Дополнительный код

В этом примере мы будем создавать разные типы оружия при помощи структуры фабрики.
*/

// Интерфейс продукта
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Класс продукта
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

// Класс конкретного продукта

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{gun{name: "ak47", power: 4}}
}

// Класс конкретного продукта

type m4 struct {
	gun
}

func newM4() iGun {
	return &m4{gun{name: "m4", power: 3}}
}

// Фабрика оружия
func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "m4":
		return newM4(), nil
	default:
		return nil, fmt.Errorf("error with type")
	}
}

func main() {
	ak47, _ := getGun("ak47")
	m4, _ := getGun("m4")

	printGun(ak47)
	printGun(m4)
}

func printGun(g iGun) {
	fmt.Println("now gun is - ", g.getName(), "and power is - ", g.getPower())
}
