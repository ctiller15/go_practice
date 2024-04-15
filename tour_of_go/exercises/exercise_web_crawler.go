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

// SafeCache is safe to use concurrently.
type SafeCache struct {
	mu    sync.Mutex
	cache map[string]bool
}

//var cache = map[string]bool{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *SafeCache) Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	c.mu.Lock()
	_, ok := c.cache[url]

	// Don't fetch same url twice.
	if ok {
		c.mu.Unlock()
		return
	} else {
		c.cache[url] = true
	}

	c.mu.Unlock()

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go c.Crawl(u, depth-1, fetcher)
	}
	return
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Starting!")
	c := SafeCache{cache: make(map[string]bool)}
	wg.Add(1)
	go c.Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()

	fmt.Println(c.cache)
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
