package main

import (
	"testing"
)

func TestRenderReturnCorrectFormatedContent(t *testing.T) {

	var readContent = `
---
title: {{title}}
location: {{location}}
link: {{link}}
publishedOn: {{publishedOn}}
type: {{type}}
---
{{content}}
`

	var args = ReportTemplate{
		"Hello World",
		"Localhost",
		"http://localhost",
		"2021-04-20T09:15:00-0400",
		"idk",
		"A wonderful test",
	}
	result, err := render(readContent, args)

	if err != nil {
		t.Errorf("An error occured in the render, this should not happen")
	}

	var expected = `
---
title: Hello World
location: Localhost
link: http://localhost
publishedOn: 2021-04-20T09:15:00-0400
type: idk
---
A wonderful test
`

	if result != expected {
		t.Errorf("The rendered text does not match expected output : got: %s =! expected: %s", result, expected)
	}
}
