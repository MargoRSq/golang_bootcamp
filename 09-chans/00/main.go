package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func sleepSort(arr []int) (ch chan int) {
	ch = make(chan int, len(arr))
	defer close(ch)
	for _, item := range arr {
		wg.Add(1)
		go func(ch chan int, item int) {
			time.Sleep(time.Duration(item) * time.Millisecond)
			ch <- item
			wg.Done()
		}(ch, item)
	}
	wg.Wait()
	return
}

func main() {

	unsorted := []int{6, 2, 1, 5}

	ch := sleepSort(unsorted)

	for num := range ch {
		fmt.Println(num)
	}

}
