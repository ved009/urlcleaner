package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

type URLParameters struct {
	URL     string
	Params  []string
	HashKey string
}

type ByHashKey []URLParameters

func (a ByHashKey) Len() int           { return len(a) }
func (a ByHashKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHashKey) Less(i, j int) bool { return a[i].HashKey < a[j].HashKey }

func main() {
	urls := readURLsFromStdin()

	parameters := make([]URLParameters, 0)

	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			fmt.Printf("Failed to parse URL: %v\n", err)
			continue
		}

		queryParams := parsedURL.Query()
		params := make([]string, 0, len(queryParams))

		for name := range queryParams {
			params = append(params, name)
		}

		sort.Strings(params)

		hashKey := strings.Join(params, "&")

		found := false
		for _, param := range parameters {
			if param.HashKey == hashKey {
				found = true
				break
			}
		}

		if !found {
			parameters = append(parameters, URLParameters{
				URL:     u,
				Params:  params,
				HashKey: hashKey,
			})
		}
	}

	sort.Sort(ByHashKey(parameters))

	// Print sorted URLs
	for _, param := range parameters {
		fmt.Println(param.URL)
	}
}

// Read URLs from standard input (text file)
func readURLsFromStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)
	urls := []string{}

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read standard input: %v\n", err)
		os.Exit(1)
	}

	return urls
}
