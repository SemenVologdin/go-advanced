package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

const (
	N = 10
)

func main() {
	numberChan, err := generateNumbers(N)
	if err != nil {
		fmt.Printf("Error while generate numbers: %s\n", err.Error())
		os.Exit(1)
	}

	numbers := make([]int, 0, N)
	for number := range squaringNumbers(numberChan) {
		numbers = append(numbers, number)
	}

	fmt.Println(numbers)
}

func generateNumbers(num int) (chan int, error) {
	if num < 0 {
		return nil, fmt.Errorf("number of elements must be positive")
	}

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for range num {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- rand.Intn(101) // [0, 100]
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch, nil
}

func squaringNumbers(numberCh chan int) chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for num := range numberCh {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- num * num
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
