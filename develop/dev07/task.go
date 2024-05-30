package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал,
если один из его составляющих каналов закроется.

Очевидным вариантом решения могло бы стать выражение при использованием select,
которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов,
с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной
функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})

	var wg sync.WaitGroup

	for i := range channels {
		wg.Add(1)
		go func(c <-chan interface{}) {
			for range c {
			}
			wg.Done()
		}(channels[i])
	}

	wg.Wait()
	close(res)

	return res
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
