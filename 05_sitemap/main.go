package main

import (
	"fmt"
	"log"
	"net/http"

	link "github.com/jcamilom/gophercises/04_link"
)

func main() {
	resp, err := http.Get("http://www.calhoun.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.Status == "200 OK" {
		links, err := link.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("HTML:\n\n", string(body))
		fmt.Printf("\n%+v", links)
	}
}
