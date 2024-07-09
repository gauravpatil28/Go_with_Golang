package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in GOlang")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{} // with the help of RW mutex we can also lock the memory which we read or write

	var digits = []int{0}

	wg.Add(4)

	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One")
		mut.Lock()
		digits = append(digits, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two")
		mut.Lock()
		digits = append(digits, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three")
		mut.Lock()
		digits = append(digits, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		mut.RLock()
		fmt.Println(digits)
		mut.RUnlock()
		w.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(digits)

}
