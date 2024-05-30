package main

import (
	"fmt"
	"os"
)

/*Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

*/

//Проверяем, что первый символ не цифра
func checkFist(runeStr []rune) error {
	if runeStr[0] >= 49 && runeStr[0] <= 57 {
		return fmt.Errorf("invalid string")
	}
	return nil
}

// Удаляем элемент, который будем распаковывать
func deleteFromSlice(sl []rune, index int) []rune {
	// Удалить элемент по индексу i из a.
	// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(sl[index:], sl[index+1:])
	// 2. Усечь срез.
	sl = sl[:len(sl)-1]
	return sl
}

// Вспомогательная функция, возвращает n символов
func addSymbols(symb rune, count int) []rune {
	var runeSl []rune
	for i := 0; i < count; i++ {
		runeSl = append(runeSl, symb)
	}
	return runeSl
}

func unpack(str string) string {
	runeStr := []rune(str)
	err := checkFist(runeStr)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	var resStr []rune
	for index, value := range runeStr {
		if index != 0 {
			if value >= 49 && value <= 57 {
				if runeStr[index-1] >= 49 && runeStr[index-1] <= 57 {
					fmt.Println("incorrect string")
					os.Exit(1)
				}
				resStr = deleteFromSlice(resStr, len(resStr)-1)
				resStr = append(resStr, addSymbols(runeStr[index-1], int(value-'0'))...)
				continue
			}

		}
		resStr = append(resStr, value)
	}
	return string(resStr)
}

func main() {
	str := `dsf4fd`
	if str == "" {
		fmt.Println("")
		return
	}
	fmt.Println(unpack(str))
}
