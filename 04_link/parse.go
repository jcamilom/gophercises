package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {

	r, err := os.Open("ex1.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	LinkParser(doc)
}

func LinkParser(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		fmt.Printf("\n%+v\n", node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		LinkParser(child)
	}
}
