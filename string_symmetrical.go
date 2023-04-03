package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "amaama"
	partStr := len(str) / 2
	if len(str)%2 != 0 {
		arrStrPart := strings.Split(str, string(str[partStr]))
		fmt.Println(strings.EqualFold(arrStrPart[0], arrStrPart[1]))
	} else {
		fmt.Println(strings.EqualFold(str[:partStr], str[partStr:]))
	}
}
