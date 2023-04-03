package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu1, mu2 sync.Mutex
		wg       sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		fmt.Println("lock1")
		defer wg.Done()

		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock()
		defer mu2.Unlock()
	}()

	go func() {
		fmt.Println("lock2")
		defer wg.Done()

		mu2.Lock()
		defer mu2.Unlock()

		mu1.Lock()
		defer mu1.Unlock()
	}()

	wg.Wait()
	fmt.Println("Successfully exited")
}
