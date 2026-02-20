package main

import (
	"strings"
	"testing"
)

func TestBuildSearchURL(t *testing.T) {

	in := Interval{start: "2012", end: "2020"}
	buildSearchURL(&in)

}

func TestBuildSearchURLWithNonSenseInterval(t *testing.T) {

	in := Interval{start: "2020", end: "2001"}
	_, err := buildSearchURL(&in)

	if err == nil {
		t.Fatalf("Expected an error but got none")
	}
	if !strings.Contains(err.Error(), "cannot be greater than ") {
		t.Errorf("Unexpected error message: %s", err.Error())
	}

}

func TestBuildSearchURLWithIncorrectInterval(t *testing.T) {
	in := Interval{start: "a", end: "2001"}
	_, err := buildSearchURL(&in)

	if err == nil {
		t.Errorf("This should raise an error")
	}
}
