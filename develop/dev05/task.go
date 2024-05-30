package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

//Чтение из файла
func readFromFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var res []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res, nil
}

//Чтение из консоли, если не передан файл
func readFromConsole() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for {
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			input = append(input, text)
		} else {
			// exit if user entered an empty string
			break
		}
	}
	return input
}

// Совпадения
func match(str string, pat string) (bool, error) {
	return regexp.MatchString(pat, str)
}

// Выводит доп строки после той, в которой нашлось совпадение
func afterPrint(lines []string, index int, after int) {
	if index+after+1 < len(lines) {
		for i := index + 1; i <= index+after; i++ {
			fmt.Println(lines[i])
		}
	} else {
		for i := index + 1; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	}
}

// Выводит строки, предшествующие совпадению
func beforePrint(lines []string, index int, before int) {
	if index-before >= 0 {
		for i := index - before; i < index; i++ {
			fmt.Println(lines[i])
		}
	} else {
		for i := 0; i < index; i++ {
			fmt.Println(lines[i])
		}
	}
}

func grep(after int, before int, count bool, invert bool, lineNum bool, pattern string, fileName string) {
	var lines []string
	var err error

	// Читаем файл, или вводи данные с клавиатуры
	if fileName == "-" {
		lines = readFromConsole()
	} else {
		lines, err = readFromFile(fileName)
	}

	if err != nil {
		log.Fatalf("cannot open file")
	}
	var cou int

	for index, val := range lines {

		ok, err := match(val, pattern)
		if err != nil {
			log.Fatalf("error with matchking")
		}

		// Инвертируем, если передан соотв. флаг
		ok = ok != invert

		if ok == true {

			cou++

			if lineNum {
				fmt.Println(index)
			} else {
				// Печатаем строки до совпадения
				if before > 0 {
					beforePrint(lines, index, before)
				}

				fmt.Println(val)

				// Печатаем строки после совпадения
				if after > 0 {
					afterPrint(lines, index, after)
				}
			}
		}
	}

	if count {
		fmt.Printf("found in %d lines\n", cou)
	}

}

func main() {
	// Считываем ключи
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	// Проверяем, что передано дополнительно два аргумента (паттер и имя файла(или -, если будет ввод с клавиатуры))
	if flag.NArg() < 2 {
		log.Fatalf("File name or - for console input and pattern required")
	}

	args := flag.Args()
	pattern := args[0]
	fileName := args[1]

	// Если передан контекст, то переопределяем в афтер и бефор
	if *context != 0 {
		after = context
		before = context
	}

	// Изменяем паттерн, если передан соотв. флаг
	if *ignoreCase {
		pattern = "(?i)" + pattern
	}

	// Редактируем паттерн на поиск целой строки
	if *fixed {
		pattern = "^" + pattern + "$"
	}

	grep(*after, *before, *count, *invert, *lineNum, pattern, fileName)
}
