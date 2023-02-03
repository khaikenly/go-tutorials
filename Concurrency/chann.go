package main

import (
	"fmt"
	"log"
	"sync"
)

func printNum(wg *sync.WaitGroup) {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d ", i)
	}
	defer wg.Done()
}

func printChar(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c ", i)
	}
	defer wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go printChar(wg)
	go printNum(wg)

	wg.Wait()
	log.Println("Done !")
}
