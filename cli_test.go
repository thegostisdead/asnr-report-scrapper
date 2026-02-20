package main

import (
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	testCases := []struct {
		rawArgs  string
		expected Config
		error    error
	}{
		{"--help", Config{}, nil},  // should display help
		{"unknown", Config{}, nil}, // should display help
		{"", Config{}, nil},        // should display help
		//
		{"--from=2000 --to=2010", Config{
			from: "2000",
			to:   "2010",
		}, nil},
		{"--from=2000", Config{from: "2000"}, nil},
		{"--to=2000", Config{to: "2000"}, nil},
		{"--from=2000 --to=2010 --out=reports", Config{from: "2000", to: "2010", out: "reports"}, nil},
		// {"--from=2000 --to=2010 --only-missing=output", Config{}, nil},

	}
	for _, testCase := range testCases {
		args := strings.Split(testCase.rawArgs, " ")
		got, err := ParseArgs(args)

		if err != testCase.error {
			t.Errorf("Command %s should an error : %s", testCase.rawArgs, testCase.error)
		}

		if got != testCase.expected {
			t.Errorf("Command %s should return a config like : %s", testCase.rawArgs, testCase.expected)
		}

	}
}
