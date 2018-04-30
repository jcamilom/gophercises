package main

import (
	"log"
	"net/http"
	"net/url"

	link "github.com/jcamilom/gophercises/04_link"
)

func main() {
	URL := "http://www.calhoun.io/"

	// The parsed URL
	parsedURL, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(URL)
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
		//fmt.Printf("\n%+v", links)

		links = filterLinks(links, parsedURL)

		/* fmt.Println("====================")
		fmt.Println(len(links), " links returned.")
		fmt.Println("====================") */

		/* for _, l := range links {
			fmt.Println("====================")
			fmt.Println("Link: ", l.Href)
			fmt.Println("Text: ", l.Text)
			fmt.Println("====================")
		} */
	}
}

func filterLinks(links []link.Link, pageURL *url.URL) []link.Link {
	// Variable where the filtered links are stored
	var ret []link.Link
	pageHostname := pageURL.Hostname()

	/* fmt.Println("====================")
	fmt.Println(len(links), " links received.")
	fmt.Println("====================") */

	for _, l := range links {
		linkURL, err := url.Parse(l.Href)
		if err != nil {
			log.Fatal(err)
		}
		linkHostname := linkURL.Hostname()
		linkPath := linkURL.EscapedPath()

		/* fmt.Println("====================")
		fmt.Println("linkHref: ", l.Href)
		fmt.Println("linkText: ", l.Text)
		fmt.Println("linkHostname: ", linkHostname)
		fmt.Println("linkPath: ", linkURL.EscapedPath())
		fmt.Println("pageHostname: ", pageHostname)
		fmt.Println("====================") */

		if linkHostname == pageHostname || (linkHostname == "" && len(linkPath) > 0) {
			//fmt.Println("Adding: ", l.Href, l.Text)
			ret = append(ret, l)
		} else {
			//fmt.Println("Removing: ", l.Href, l.Text)
		}
	}
	return ret
}
