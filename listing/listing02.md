Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Вывод программы:
2
1
В неименованной функции anotherTest происходит копирование значения x
в место возврата, после чего выполняется defer функция. В test же при 
вызове return копирование не происходит и defer изменяет ту же самую переменную,
которая и возвращается.
```
