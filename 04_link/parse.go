package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Link holds the links struct
type Link struct {
	Href string
	Text string
}

func main() {

	r, err := os.Open("ex2.html")
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

// LinkParser searchs for links inside the three of the passed node
func LinkParser(node *html.Node, links *[]Link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		// Extract the text from the link
		text := extractTextFromLink(node)
		// Get the link
		for _, att := range node.Attr {
			// Gets the attribute "href" and stores it in the slice
			if att.Key == "href" {
				*links = append(*links, Link{Href: att.Val, Text: text})
			}
		}
		//fmt.Printf("==\n%+v\n==\n", node)
	} else {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			LinkParser(child, links)
		}
	}
	//fmt.Printf("==\n%v ---> %+v\n==\n", &node, node)
}

// Extracts the text inside the <a></a> element and returns it
func extractTextFromLink(node *html.Node) (text string) {
	for linkChild := node.FirstChild; linkChild != nil; linkChild = linkChild.NextSibling {
		if linkChild.Parent.Data == "a" {
			fmt.Printf("==\n%+v\n==\nParentData: %v\n", linkChild, linkChild.Parent.Attr[0].Val)
		}
		if trimedData := strings.TrimSpace(linkChild.Data); linkChild.Type == html.TextNode && len(trimedData) > 0 {
			//fmt.Printf("\n\ntext=%v, len=%v\n\n", trimedData, len(trimedData))
			if text == "" {
				text = trimedData
			} else {
				text += " " + trimedData
			}
			//fmt.Println(text)
		}
		if text2 := extractTextFromLink(linkChild); text2 != "" {
			if text == "" {
				text = text2
			} else {
				text += " " + text2
			}
		}
	}
	return
}
