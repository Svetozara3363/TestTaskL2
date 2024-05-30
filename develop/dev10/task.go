package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа
должна также завершаться. При подключении к несуществующему серверу, программа должна завершаться через timeout

*/

func connect(ip string, port string) {
	// Подключаемся к сокету
	address := ip + ":" + port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("connection failed")
	}
	fmt.Println("Connection start")
	for {

		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if message == "" {
			fmt.Println("connect close")
			break
		}

		fmt.Print("Message from server: " + message)

	}
}

func main() {
	// Считываем задержку на подключение
	timeout := flag.Int("t", 10, "timeout")

	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatalln("Invalid input(port or ip)")
	}

	args := flag.Args()

	ip := args[0]
	port := args[1]

	fmt.Println("Wait to connection")
	time.Sleep(time.Second * time.Duration(*timeout))

	connect(ip, port)
}
