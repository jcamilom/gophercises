package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jcamilom/gophercises/04_link"
)

func main() {
	r, err := os.Open("ex2.html")
	if err != nil {
		log.Fatal(err)
	}

	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
