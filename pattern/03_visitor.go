package main

import "fmt"

/*
Посетитель — это поведенческий паттерн,
который позволяет добавить новую операцию для целой иерархии классов,
не изменяя код этих классов.

Плюсы:
-Нет необходимости изменять классы
-Похожие операции над разными объектами хранятся в одном месте

Минусы:
-Лишний код
-Может привести к нарушению инкапсуляции элементов.
*/

// Общий интерфейс посетителя
type visitor interface {
	visitForSquare(*square)
	visitForTriangle(*triangle)
	visitForCircle(*circle)
}

//----------------------------------------------------------------------------------------------

// Конкретный визитор, в данном случае, для подсчета площади фигур
//----------------------------------------------------------------------------------------------
type areaCalcVisitor struct {
	area int
}

func (a *areaCalcVisitor) visitForSquare(s *square) {
	a.area = s.side * s.side
}

func (a *areaCalcVisitor) visitForTriangle(t *triangle) {
	a.area = (t.a * t.b) / 2
}

func (a *areaCalcVisitor) visitForCircle(c *circle) {
	a.area = int(float64(c.radius) * float64(c.radius) * 3.14)
}

//----------------------------------------------------------------------------------------------

// Общий интерфейс для всех фигур в который мы добавили метод accept
type shape interface {
	getType() string
	accept(visitor)
}

// Конкретный класс фигуры
//----------------------------------------------------------------------------------------------
type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

//----------------------------------------------------------------------------------------------

// Конкретный класс фигуры
//----------------------------------------------------------------------------------------------
type triangle struct {
	a, b, c int
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}

func (t *triangle) getType() string {
	return "Triangle"
}

//----------------------------------------------------------------------------------------------

// Конкретный класс фигуры
//----------------------------------------------------------------------------------------------
type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

//----------------------------------------------------------------------------------------------

func main() {
	// Создаем объекты фигур
	square := &square{side: 5}
	triangle := &triangle{a: 3,
		b: 4,
		c: 5}
	circle := &circle{radius: 5}

	areaCalcVisitor := &areaCalcVisitor{}
	square.accept(areaCalcVisitor)
	fmt.Println("area square is -", areaCalcVisitor.area)
	triangle.accept(areaCalcVisitor)
	fmt.Println("area triangle is -", areaCalcVisitor.area)
	circle.accept(areaCalcVisitor)
	fmt.Println("area circle is -", areaCalcVisitor.area)

}
