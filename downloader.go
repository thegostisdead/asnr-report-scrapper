package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// https://gist.github.com/nevermosby/b54d473ea9153bb75eebd14d8d816544
// TODO read this https://abhii85.hashnode.dev/understanding-concurrency-in-go-and-building-a-concurrent-file-downloader
func DownloadExtractSave(urls []url.URL, outputDir string) error {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, u := range urls {
		go func(u string) {
			defer wg.Done()

			fmt.Println("Downloading", u)

			res, err := http.Get(u)
			if err != nil {
				log.Fatal("http get error: ", err)
				return "qd"
			}

			if err != nil {
				log.Fatal("Error while downloading", u, "-", err)
				return "qd"
			}

			fmt.Println("Downloaded ", u)

			go htmlToTemplate(res.Body) // directly processing the file, its IO bound

		}(u.String())
	}

	wg.Wait()
	fmt.Println("Done")

	return nil

}
