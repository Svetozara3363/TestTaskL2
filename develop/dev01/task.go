package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*Создать программу печатающую точное время с использованием NTP -библиотеки.
Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS
*/

// Функция, которая возвращает текущее системное время
func osTime() string {
	return time.Now().Format("15:04:05.999")
}

// Функция, которая возвращает текущее сетевое время
func netTime() (string, error) {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return ntpTime.Format("15:04:05.999"), err
}

// Функция, которая возвращает текущее системное время с учетом задержки
func netTimeWithDuration() (string, error) {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	dTime := time.Now().Add(response.ClockOffset)
	return dTime.Format("15:04:05.999"), err
}

func main() {
	nowTime := osTime()
	fmt.Println("Now system time is - ", nowTime)
	nowTime, err := netTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Println("Net time is - ", nowTime)
	nowTime, err = netTimeWithDuration()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Println("Net time with duration is - ", nowTime)
}
