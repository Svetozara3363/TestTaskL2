package main

import "fmt"

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Строитель позволяет не использовать длинные конструкторы.

Если потребуется объединить шаги создания создается объект директор, определяющий
порядок шагов, так процесс создания будет полностью скрыт от клиента
(директор - не является обязательным в реализации паттерна)

Плюсы:
- Создает объекты пошагово (если это требуется, например, деревья)
- Позволяет создавать несколько представлений одного объекта, переиспользуя код
Минусы:
- Усложняет код, требуется написание дополнительных классов, может усложняться внедрнее зависимостей

*/

// Класс конструируемого продукта
type car struct {
	brand   string
	engine  float32
	premium bool
}

const (
	ladaBuilderType    = "lada"
	ferrariBuilderType = "ferrari"
)

// Общий интерфейс для строителей
//------------------------------------------------------------------------------------------
type builder interface {
	setBrand()
	setEngine()
	setPremium()
	getCar() car
}

// Выбор подходящего строителя
func getBuilder(builderType string) builder {
	switch builderType {
	case ladaBuilderType:
		return &ladaBuilder{}
	case ferrariBuilderType:
		return &ferrariBuilder{}
	default:
		return nil
	}
}

//------------------------------------------------------------------------------------------

// Реализация конкретного строителя для Лада
//------------------------------------------------------------------------------------------
type ladaBuilder struct {
	brand   string
	engine  float32
	premium bool
}

func newLadaBuilder() *ladaBuilder {
	return &ladaBuilder{}
}

func (b *ladaBuilder) setBrand() {
	b.brand = "Lada"
}

func (b *ladaBuilder) setEngine() {
	b.engine = 1.6
}

func (b *ladaBuilder) setPremium() {
	b.premium = false
}

func (b *ladaBuilder) getCar() car {
	return car{
		brand:   b.brand,
		engine:  b.engine,
		premium: b.premium,
	}
}

//------------------------------------------------------------------------------------------

// Реализация конкретного строителя для Феррари
//------------------------------------------------------------------------------------------
type ferrariBuilder struct {
	brand   string
	engine  float32
	premium bool
}

func newFerrariBuilder() *ferrariBuilder {
	return &ferrariBuilder{}
}

func (b *ferrariBuilder) setBrand() {
	b.brand = "Ferrari"
}

func (b *ferrariBuilder) setEngine() {
	b.engine = 6.0
}

func (b *ferrariBuilder) setPremium() {
	b.premium = true
}

func (b *ferrariBuilder) getCar() car {
	return car{
		brand:   b.brand,
		engine:  b.engine,
		premium: b.premium,
	}
}

//------------------------------------------------------------------------------------------

func main() {
	ladaCarBuilder := getBuilder("lada")
	ferrariCarBuilder := getBuilder("ferrari")

	ladaCarBuilder.setEngine()
	ladaCarBuilder.setBrand()
	ladaCarBuilder.setPremium()
	car := ladaCarBuilder.getCar()
	fmt.Printf("%s has engine - %1.1f, Premium - %v\n", car.brand, car.engine, car.premium)

	ferrariCarBuilder.setEngine()
	ferrariCarBuilder.setBrand()
	ferrariCarBuilder.setPremium()
	car = ferrariCarBuilder.getCar()
	fmt.Printf("%s has engine - %1.1f, Premium - %v\n", car.brand, car.engine, car.premium)
}
