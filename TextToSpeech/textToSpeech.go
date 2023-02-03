package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	htgotts "github.com/hegedustibor/htgo-tts"
)

func main() {
	err := os.RemoveAll("./audio/") // delete an entire directory
	if err != nil {
		fmt.Println(err)
	}
	rand.Seed(time.Now().Unix())
	num := rand.Intn(99)
	fmt.Println(num)
	speech := htgotts.Speech{Folder: "audio", Language: "vi"}
	speech.Speak("số " + fmt.Sprint(num))
}
