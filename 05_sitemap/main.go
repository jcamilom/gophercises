package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	link "github.com/jcamilom/gophercises/04_link"
)

func main() {
	urlFlag := flag.String("url", "https://www.calhoun.io", "the url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.Status == "200 OK" {
		reqUrl := resp.Request.URL
		baseUrl := &url.URL{
			Scheme: reqUrl.Scheme,
			Host:   reqUrl.Host,
		}
		base := baseUrl.String()

		links, err := link.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("HTML:\n\n", string(body))
		//fmt.Printf("\n%+v", links)

		var hrefs []string
		for _, l := range links {
			switch {
			case strings.HasPrefix(l.Href, "/"):
				hrefs = append(hrefs, base+l.Href)
			case strings.HasPrefix(l.Href, "http"):
				hrefs = append(hrefs, l.Href)
			}
		}

		for _, href := range hrefs {
			fmt.Println(href)
		}

		//links = filterLinks(links, parsedURL)

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
