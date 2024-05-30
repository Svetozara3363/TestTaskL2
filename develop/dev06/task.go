package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

*/

func cut(fields string, delimiter string, separated bool) {
	scanner := bufio.NewScanner(os.Stdin)
	// Парсим столбцы, которые просят вывести
	field := strings.Split(fields, ",")
	var fieldsP []int
	for _, v := range field {
		strInt, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("error fields input")
		}
		fieldsP = append(fieldsP, strInt)
	}

	// Бесконечный цикл для ввода текста
	for {
		ok := scanner.Scan()
		if !ok && scanner.Err() == nil {
			break
		}
		str := scanner.Text()

		if separated {
			if !strings.Contains(str, delimiter) {
				continue
			}
		}

		strD := strings.Split(str, delimiter)

		for _, v := range fieldsP {
			fmt.Print(strD[v], " ")
		}
		fmt.Println()
	}
}

func main() {

	fields := flag.String("f", "0,2", "fields")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", true, "separated")

	cut(*fields, *delimiter, *separated)

}
