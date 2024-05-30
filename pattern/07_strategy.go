package main

import "fmt"

/*
Стратегия — это поведенческий паттерн проектирования,
который определяет семейство схожих алгоритмов и помещает каждый из них
в собственный класс, после чего алгоритмы можно взаимозаменять прямо во
время исполнения программы.

Другие объекты содержат ссылку на объект-стратегию и делегируют ей работу.
Программа может подменить этот объект другим, если требуется иной способ решения задачи.

Плюсы:
- Горячая замена алгоритмов на лету.
- Изолирует код и данные алгоритмов от остальных классов.
- Уход от наследования к делегированию.

Минусы:
-  Усложняет программу за счёт дополнительных классов.
*/

// Интерфейс стратегии
type evictionAlgo interface {
	evict(c *cache)
}

//-----------------------------------------------------------------------------------------------------------------

// Конкретная стратегия
type fifo struct {
}

func (f *fifo) evict(c *cache) {
	fmt.Println("now we use fifo")
}

//-----------------------------------------------------------------------------------------------------------------

// Вторая стратегия
type lru struct {
}

func (f *lru) evict(c *cache) {
	fmt.Println("now we use lru")
}

//-----------------------------------------------------------------------------------------------------------------

// Третья стратегия
type lfu struct {
}

func (f *lfu) evict(c *cache) {
	fmt.Println("now we use lfu")
}

//-----------------------------------------------------------------------------------------------------------------

// Объект кэша
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

//-----------------------------------------------------------------------------------------------------------------

func main() {

	// Создаем объекты стратегий
	fifo := &fifo{}
	lru := &lru{}
	lfu := &lfu{}

	// Создаем объект кеша
	cache := initCache(fifo)

	cache.add("a", "1")
	cache.add("b", "2")

	// На третьем вызове используется указанная раннее стратегия для очистки кеша
	cache.add("c", "3")

	// Меняем стратегию очистки кеша
	cache.setEvictionAlgo(lru)

	// Опять использукем метод очистки
	cache.add("d", "4")

	// Снова меняем и используем стратегию
	cache.setEvictionAlgo(lfu)
	cache.add("e", "5")
}
