package main

import (
	"fmt"
	"sync"
	"time"
)

type Url string

func (url Url) string() string {
	return string(url)
}

type Urls struct {
	// Urls is a thread-safe collection of urls
	v   map[Url]int
	mux sync.Mutex
}

func (p *Urls) add(url Url) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.v[url]++
}

func (p *Urls) list() []Url {
	p.mux.Lock()
	defer p.mux.Unlock()
	urlSlice := make([]Url, 0, len(p.v))
	for k := range p.v {
		urlSlice = append(urlSlice, k)
	}
	return urlSlice
}

func (p *Urls) get(k Url) int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.v[k]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, collection *Urls) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	collection.add(Url(url))

	if collection.v[Url(url)] == 1 {
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go Crawl(u, depth-1, fetcher, collection)
		}
	}

	return
}

func main() {
	collection := &Urls{v: make(map[Url]int)}
	go Crawl("https://golang.org/", 4, fetcher, collection)
	time.Sleep(time.Second)
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
