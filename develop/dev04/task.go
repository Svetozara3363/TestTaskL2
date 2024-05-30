package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

// Перевод всех символов в нижний регистр
func toLower(words []string) []string {
	var res []string
	for _, val := range words {
		res = append(res, strings.ToLower(val))
	}
	return res
}

// Удаление повторяющихся слов
func deleteRepeat(words []string) []string {
	var res []string
	m := make(map[string]bool)

	for _, v := range words {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func searchForAnagrams(words []string) map[string][]string {
	// Создаем результирующую мапу
	mapH := make(map[string][]string)

	// Заранее переводим все символы в нижний регистр и удаляем повторения
	wordsLower := toLower(words)
	finWords := deleteRepeat(wordsLower)

	for _, val := range finWords {
		word := []rune(val)
		sort.Slice(word, func(i, j int) bool {
			return word[j] > word[i]
		})

		wordString := string(word)
		mapH[wordString] = append(mapH[wordString], val)
	}
	resMap := make(map[string][]string)
	for _, val := range mapH {
		if len(val) > 1 {

			resMap[val[0]] = val
			sort.Strings(val)
		}
	}
	return resMap
}

func main() {
	mas := []string{"пятка", "слиток", "пятак", "черешня", "тяпка", "столик"}
	res := searchForAnagrams(mas)
	fmt.Println(res)
}
