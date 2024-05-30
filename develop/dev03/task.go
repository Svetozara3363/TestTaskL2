package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры): на входе подается
файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

*/

func sortByColumn(lines []string, ColumnNumber int, sortNumber bool) []string {
	// Создаем мапу, где ключи - слова в строках на k-ой позиции, значения - строки, содержащие в себе ключ
	m := map[string]string{}
	// Слайс ключей для сортировки
	var keys []string
	// Результирующий слайс
	var res []string
	// Проходим по слайсу строк, каждую строку разделяем по пробелам на слова,
	// слово на заданной позиции помещаем в слайс ключей для сортировки
	// добавляем в мапу нужные значения
	for i := 0; i < len(lines); i++ {
		arr := strings.Split(lines[i], " ")
		keys = append(keys, arr[ColumnNumber])
		m[arr[ColumnNumber]] = lines[i]
	}
	fmt.Println(keys)
	// Если ключ задан, то сортируем ключи по числовому значению в столбце, иначе сортируем по словам
	if sortNumber {
		keys = sortByNumber(keys)
	} else {
		sort.Strings(keys)
	}
	fmt.Println(keys)
	for _, v := range keys {
		res = append(res, m[v])
	}
	fmt.Println(res)
	return res
}

// Функция для сортировки чисел по их значению
func sortByNumber(lines []string) []string {
	var buf []int
	var res []string
	// Переводим строки в слайс чисел
	for i := 0; i < len(lines); i++ {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatal(err)
		}
		buf = append(buf, num)
	}

	sort.Ints(buf)

	for i := 0; i < len(buf); i++ {
		res = append(res, strconv.Itoa(buf[i]))
	}
	return res

}

// Функция для удаления одинаковых строк
func deleteSimilar(lines []string) []string {
	m, uniq := make(map[string]struct{}), make([]string, 0, len(lines))
	for _, v := range lines {
		if _, ok := m[v]; !ok {
			m[v], uniq = struct{}{}, append(uniq, v)
		}
	}
	return uniq
}

// Функция для переворота слайса строк
func reverseSlice(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// Функция чтения из файла
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

func sortStrings(file string, ColumnNumber int, sortNumber bool, sortReverse bool, sortUnique bool) {
	// Читаем с файла, помещаем результат в слайс строк
	lines, err := readFromFile(file)
	if err != nil {
		fmt.Println("err - ", err)
		os.Exit(1)
	}

	//Если передан ключ на сортировку строк по опрределенному столбцу
	if ColumnNumber >= 0 {
		lines = sortByColumn(lines, ColumnNumber, sortNumber)
	} else {
		sort.Strings(lines)
		if sortNumber {
			lines = sortByNumber(lines)
		}
	}

	// Реверсируем, если передан соответствующий ключ
	if sortReverse {
		reverseSlice(lines)
	}

	// Удаляем одинаковые строки
	if sortUnique {
		lines = deleteSimilar(lines)
	}

	// Выводим результат в консоль
	for _, val := range lines {
		fmt.Println(val)
	}

}

func main() {
	//Задаем и считываем ключи при запуске программы
	file := flag.String("f", "test.txt", "File")
	ColumnNumber := flag.Int("k", -1, "Sorting by Columns")
	sortNumber := flag.Bool("n", false, "Sorting by numbers")
	sortReverse := flag.Bool("r", false, "Revers sort")
	sortUnique := flag.Bool("u", false, "Unique sort")

	flag.Parse()

	sortStrings(*file, *ColumnNumber, *sortNumber, *sortReverse, *sortUnique)
}
