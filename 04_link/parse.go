package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	dfs(doc, "")

	//links := make([]Link, 0, 5) // Slice to store the links

	//LinkParser(doc, &links)

	//fmt.Printf("\n%+v\n", links)
	return nil, nil
}

func dfs(node *html.Node, padding string) {
	msg := node.Data
	if node.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}

// LinkParser searchs for links inside the three of the passed node.
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

// extractTextFromLink extracts the text nested in the HTML element and returns it.
func extractTextFromLink(node *html.Node) (text string) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		/* if child.Parent.Data == "a" {
			fmt.Printf("==\n%+v\n==\nParentData: %v\n", child, child.Parent.Attr[0].Val)
		} */
		if child.Type == html.TextNode {
			// Calls the join function if there is something to join
			if trimedData := strings.TrimSpace(child.Data); len(trimedData) > 0 {
				//fmt.Printf("\n\ntext=%v, len=%v\n\n", trimedData, len(trimedData))
				text = joinStrings(text, trimedData)
				//fmt.Println(text)
			}
		} else if textFromChilds := extractTextFromLink(child); textFromChilds != "" {
			// This previous "if" calls the join function if there is something to join
			text = joinStrings(text, textFromChilds)
		}
	}
	return
}

// joinString joins two strings and place a whitespace in between, when needed.
func joinStrings(a, b string) string {
	if a == "" {
		return b
	}
	return a + " " + b
}
