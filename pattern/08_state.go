package main

import (
	"fmt"
	"log"
)

/*
Состояние — это поведенческий паттерн проектирования, который позволяет
объектам менять поведение в зависимости от своего состояния. Извне создаётся
впечатление, что изменился класс объекта.

Основная идея в том, что программа может находиться в одном из нескольких
состояний, которые всё время сменяют друг друга. Набор этих состояний,
а также переходов между ними, предопределён и конечен. Находясь в разных
состояниях, программа может по-разному реагировать на одни и те же события,
которые происходят с ней.

Плюсы:
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.

Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

Пример сделаем на основе торгового автомата, который может пребывать только
в одном из 4 состояний:

-hasItem (имеетПредмет)
-noItem (неИмеетПредмет)
-itemRequested (выдаётПредмет)
-hasMoney (получилДеньги)

И может выполнять только 4 действия:

-Выбрать предмет
-Добавить предмет
-Ввести деньги
-Выдать предмет
*/

// Класс торгового автомата
type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{vendingMachine: v}
	itemRequestedState := &itemRequestedState{vendingMachine: v}
	hasMoneyState := &hasMoneyState{vendingMachine: v}
	noItemState := &noItemState{vendingMachine: v}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d item", count)
	v.itemCount += count
}

//------------------------------------------------------------------------

//Интерфейс состояний
type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

//Конкретное состояние
type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) insertMoney(money int) error {
	return fmt.Errorf("no item")
}

func (n *noItemState) dispenseItem() error {
	return fmt.Errorf("no item")
}

func (n *noItemState) requestItem() error {
	return fmt.Errorf("no item")
}

func (n *noItemState) addItem(count int) error {
	n.vendingMachine.incrementItemCount(count)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

//Конкретное состояние
type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h hasItemState) addItem(count int) error {
	fmt.Printf("%d items added", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h hasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("no item to request")
	}
	fmt.Println("Item requested")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h hasItemState) insertMoney(money int) error {
	return fmt.Errorf("select item first")
}

func (h hasItemState) dispenseItem() error {
	return fmt.Errorf("select item first")
}

//Конкретное состояние
type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i itemRequestedState) addItem(i2 int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("money ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i itemRequestedState) dispenseItem() error {
	return fmt.Errorf("enter money")
}

//Конкретное состояние
type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h hasMoneyState) addItem(i int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (h hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (h hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (h hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

}
