package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetched SynchMap

type SynchMap struct {
	m   map[string]bool
	mux sync.Mutex
}

func fetch(url string, fetcher Fetcher, urlchan chan []string) {
	_, ok := fetched.m[url]
	if !ok {
		body, urls, err := fetcher.Fetch(url)

		fetched.mux.Lock()
		fetched.m[url] = true
		fetched.mux.Unlock()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("found: %s %q\n", url, body)
		}

		urlchan <- urls
	}
}

// Invokes Goroutines for each url in a slice
// Returns the number of invoked Goroutines
func invoke(urls []string, fetcher Fetcher, c chan []string) int {
	count := 0
	// Invoke any expected urls
	for _, nextUrl := range urls {
		_, ok := fetched.m[nextUrl]
		if !ok {
			count++
			go fetch(nextUrl, fetcher, c)
		}
	}
	return count
}

func results(count int, c chan []string) []string {
	// Empty the existing array
	next := make([]string, 0)
	// For each invoked Goroutine
	for i := count; i > 0; i-- {
		x := <-c
		next = append(next, x...)
	}
	return next
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	if depth <= 0 {
		return
	}

	c := make(chan []string)

	next := []string{url}
	for len(next) > 0 {
		count := invoke(next, fetcher, c)
		next = results(count, c)
	}
	return
}

func main() {
	fetched.m = make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
