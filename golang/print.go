package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	letter, number := make(chan int), make(chan int)

	go func() {
		for {
			select {
			case i, ok := <-number:
				if !ok {
					close(letter)
					wg.Done()
					return
				}
				fmt.Printf("%d%d", i, i+1)
				letter <- i
			}
		}
	}()

	go func() {
		for {
			select {
			case i, ok := <-letter:
				if !ok || i >= 26 {
					close(number)
					wg.Done()
					return
				}
				fmt.Printf("%c%c", 'A'+i-1, 'A'+i)
				number <- i + 2
			}
		}
	}()

	number <- 1
	wg.Wait()

	fmt.Println()
}
