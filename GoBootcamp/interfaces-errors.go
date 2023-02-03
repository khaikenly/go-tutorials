package main

import (
	"fmt"
	"time"
)

type error interface {
	Errors() string
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Errors() (err string) {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// func Run() error {
// 	return &MyError{time.Now(), "it didn't work"}
// }

func main() {
	err := &MyError{
		When: time.Now(),
		What: "it didn't work",
	}

	if err != nil {
		fmt.Println(err.Errors())
	}
}
