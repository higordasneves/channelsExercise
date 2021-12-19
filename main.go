package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	go send(ch)
	receive(ch)
}

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 1; j <= 10; j++ {
				ch <- (i * 10) + j
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(ch)
}

func receive(ch chan int) {
	s := make([]int, 0, 100)
	for v := range ch {
		s = append(s, v)
	}
	sort.Ints(s)
	fmt.Println(s, "size:", len(s))
	fmt.Println("change on test")
}
