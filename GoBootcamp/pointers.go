package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	res, err := client.Get("http://gobootcamp.com")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res.StatusCode)
}
