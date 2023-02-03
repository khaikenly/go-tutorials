package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorestful/chapter1/romannumerals"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		if len(urlPathElements) < 3 {
			for _, v := range romannumerals.Numerals {
				fmt.Fprintf(w, "%q\n",
					html.EscapeString(v))
			}
			return
		}

		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// If resource is not in the list, send Not Found status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q",
					html.EscapeString(romannumerals.Numerals[number]))
			}
		} else {
			// For all other requests, tell that Client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
