package main

import (
	"fmt"
	"os"
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

}
