package main

import "fmt"

/*
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли
передавать запрос дальше по цепи.

Как и многие другие поведенческие паттерны, Цепочка обязанностей базируется на том,
чтобы превратить отдельные поведения в объекты. В нашем случае каждая проверка переедет в
отдельный класс с единственным методом выполнения. Данные запроса, над которым происходит проверка,
будут передаваться в метод как аргументы.

Плюсы:
-Уменьшает зависимость между клиентом и обработчиками.
-Реализует принцип единственной обязанности.

Минусы:
Запрос может остаться никем не обработанным.
*/

// Пациент
type patient struct {
	name         string
	registration bool
	doctorCheck  bool
	medicine     bool
	payment      bool
}

// Интерфейс обработчика событий
type departament interface {
	execute(*patient)
	setNext(departament)
}

//----------------------------------------------------------------------------------------------------------

// Обработчик регистратуры
type reception struct {
	next departament
}

func (r *reception) execute(p *patient) {
	if p.registration {
		fmt.Println("this patient is already registered")
		r.next.execute(p)
		return
	}
	p.registration = true
	fmt.Println("registration is completed")
	r.next.execute(p)
}

func (r *reception) setNext(next departament) {
	r.next = next
}

//----------------------------------------------------------------------------------------------------------

type doctor struct {
	next departament
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheck {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheck = true
	d.next.execute(p)
}

func (d *doctor) setNext(next departament) {
	d.next = next
}

//----------------------------------------------------------------------------------------------------------

type medical struct {
	next departament
}

func (m *medical) execute(p *patient) {
	if p.medicine {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	p.medicine = true
	fmt.Println("Medical giving medicine to patient")
	m.next.execute(p)
}

func (m *medical) setNext(next departament) {
	m.next = next
}

//----------------------------------------------------------------------------------------------------------

type cashier struct {
	next departament
}

func (c *cashier) execute(p *patient) {
	if p.medicine {
		fmt.Println("Payment Done")
		return
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next departament) {
	c.next = next
}

func main() {
	// Поочередно создаем объекты обработчики и указываем на следующие
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	//Создаем пациента Рудольфа
	patient := &patient{name: "Rudolf"}

	//Запускаем Рудольфа в приемное отделение
	reception.execute(patient)
}
