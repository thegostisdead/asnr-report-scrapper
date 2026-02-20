package main

import (
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	from string
	to   string
	out  string
	// onlyMissing string
}

func printUsage() {
	fmt.Println("Usage of scrapper :")
	fmt.Println("./scrapper --help")
	fmt.Println("./scrapper --from=2000 --to=2010")
	fmt.Println("./scrapper --from=2000")
	fmt.Println("./scrapper --to=2000")
	fmt.Println("./scrapper --from=2000 --to=2010 --out reports")
	// fmt.Println("./scrapper --from=2000 --to=2010 --only-missing output")
}

func ParseArgs(argsWithoutProg []string) (Config, error) {

	if len(argsWithoutProg) == 0 || argsWithoutProg[0] == "--help" || !strings.HasPrefix(argsWithoutProg[0], "--from") && !strings.HasPrefix(argsWithoutProg[0], "--to") {
		printUsage()
		return Config{}, nil
	}

	fs := flag.NewFlagSet("scrapper", flag.ContinueOnError)
	from := fs.String("from", "", "Start year")
	to := fs.String("to", "", "End year")
	out := fs.String("out", "", "Output directory")

	err := fs.Parse(argsWithoutProg)
	if err != nil {
		return Config{}, err
	}

	return Config{
		from: *from,
		to:   *to,
		out:  *out,
	}, nil
}
