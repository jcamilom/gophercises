package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	link "github.com/jcamilom/gophercises/04_link"
)

func main() {
	urlFlag := flag.String("url", "https://www.calhoun.io", "the url that you want to build a sitemap for")
	flag.Parse()

	pages := get(*urlFlag)

	for _, page := range pages {
		fmt.Println(page)
	}

	//links = filterLinks(links, parsedURL)
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, l := range links {
		if keepFn(l) {
			ret = append(ret, l)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
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
