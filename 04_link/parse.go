package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

// Link holds the links struct
type Link struct {
	Href string
	Text string
}

func main() {

	r, err := os.Open("ex1.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	links := make([]Link, 0, 5) // Slice to store the links

	LinkParser(doc, &links)

	fmt.Printf("\n%+v\n", links)

}

// LinkParser prints out the links in the three supplied in the html.Node
func LinkParser(node *html.Node, links *[]Link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, att := range node.Attr {
			*links = append(*links, Link{Href: att.Val, Text: att.Namespace})
		}
		//fmt.Printf("att len=%v\n", len(node.Attr))
		//fmt.Printf("\n%+v\n", node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		LinkParser(child, links)
	}
}
