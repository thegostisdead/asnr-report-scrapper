package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	ae := os.Args[1:]

	config, err := ParseArgs(ae)

	if err != nil {
		fmt.Println("Error parsing arguments")
	}

	fmt.Println(config)

	var interval Interval
	if config.from == "" && config.to == "" {
		fmt.Println("Auto detecting interval...")
		interval = DetectAvailableInterval()

	} else {
		interval = Interval{
			start: config.from,
			end:   config.to,
		}
	}

	fmt.Println("Interval will be :", interval)

	urls, err := searchWithRange(&interval)

	output := config.out

	if output == "" {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		output = filepath.Dir(ex)
	}

	err := DownloadExtractSave(urls, config.out)

	if err != nil {
		//
	}

}
