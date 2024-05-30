package main

import (
	"fmt"
	"log"
)

/*
Паттерн «Фасад» предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы.
Фасад определяет интерфейс более высокого уровня, который упрощает использование подсистемы.
*/


// Класс кошелька
type wallet struct {
	balance float64
}

// Конструктор кошелька
func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

// Метод добавления денег в кошелек
func (w *wallet) addMoney(money float64) {
	w.balance += money
	fmt.Println("money add!")
}


// Класс Юзера
type user struct {
	name string
}

// Конструктор Юзера
func newUser(name string) *user {
	return &user{
		name: name,
	}
}

// Метод проверки Юзера
func (user *user) checkUser(userName string) error {
	if user.name != userName {
		return fmt.Errorf("user name is incorrect")
	}
	fmt.Println("User is verified")
	return nil
}


// Класс фасада
type walletFacade struct {
	wallet *wallet
	user   *user
}

//Конструктор фасада
func newWalletFacade(user string) *walletFacade {
	return &walletFacade{
		wallet: newWallet(),
		user:   newUser(user),
	}
}

// Метод фасада по проверке юзера и добавлении ему денег
func (facade *walletFacade) addMoney(user string, money float64) error {
	err := facade.user.checkUser(user)
	if err != nil {
		return err
	}
	facade.wallet.addMoney(money)
	fmt.Println("Facade done")
	return nil
}


func main() {
	facade := newWalletFacade("Super")
	err := facade.addMoney("Super", 100)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("Current balance:", facade.wallet.balance)
}
