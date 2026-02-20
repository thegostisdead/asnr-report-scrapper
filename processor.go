package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

var baseUrl = "https://reglementation-controle.asnr.fr/controle/actualites-du-controle/installations-nucleaires/avis-d-incident-des-installations-nucleaires"
var filtersRegex = regexp.MustCompile(`(?mi)window\.filters\s=\sJSON\.parse\('(.*?)\'\);`)

type Interval struct {
	start string
	end   string
}

type Filter struct {
	Name    string      `json:"name"`
	Entries interface{} `json:"entries"`
	Label   string      `json:"label"`
}

func DetectAvailableInterval() Interval {
	res, err := http.Get(baseUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error parsing the request body: %s\n", err)
		os.Exit(1)
	}

	matches := filtersRegex.FindStringSubmatch(string(resBody))
	if len(matches) < 2 {
		fmt.Printf("No matches found\n")
		os.Exit(1)
	}

	capturedGroup := matches[1]

	unescaped, err := _UnescapeUnicodeCharactersInJSON(capturedGroup)
	if err != nil {
		fmt.Printf("Error unescaping: %s\n", err)
		os.Exit(1)
	}

	var filters []Filter
	err = json.Unmarshal(unescaped, &filters)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %s\n", err)
		os.Exit(1)
	}

	// Process the filters
	for _, filter := range filters {
		fmt.Printf("\n=== %s ===\n", filter.Label)

		// Check if entries is a map (most common case)
		if entriesMap, ok := filter.Entries.(map[string]interface{}); ok {
			for key, value := range entriesMap {
				fmt.Printf("  %s: %v\n", key, value)
			}
		} else if entriesArray, ok := filter.Entries.([]interface{}); ok {
			// Handle empty arrays
			if len(entriesArray) == 0 {
				fmt.Printf("  (empty)\n")
			}
		}
	}

	return Interval{}
}

func buildSearchURL(interval *Interval) (string, error) {
	base := "https://reglementation-controle.asnr.fr/controle/actualites-du-controle/installations-nucleaires/avis-d-incident-des-installations-nucleaires"

	startAsInt, err := strconv.Atoi(interval.start)
	if err != nil {
		return "", fmt.Errorf("invalid start value: %w", err)
	}

	endAsInt, err := strconv.Atoi(interval.end)
	if err != nil {
		return "", fmt.Errorf("invalid end value: %w", err)
	}

	if startAsInt > endAsInt {
		return "", fmt.Errorf("start year %d cannot be greater than end year %d", startAsInt, endAsInt)
	}

	params := url.Values{}
	if interval.start != "" {
		params.Set("publication_date_year[from]", interval.start)
	}
	if interval.end != "" {
		params.Set("publication_date_year[to]", interval.end)
	}

	return base + "?" + params.Encode(), nil

}

func searchWithRange(interval *Interval) ([]url.URL, error) {
	toFetch, err := buildSearchURL(interval)

	if err != nil {
		panic(err)
	}
	fmt.Println(toFetch)
	return nil, nil

}
