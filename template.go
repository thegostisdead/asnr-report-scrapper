package main

import "strings"

type ReportTemplate struct {
	Title     string
	Location  string
	Link      string
	Published string
	Type      string
	Content   string
}

func render(templateContent string, args ReportTemplate) (string, error) {

	replacer := strings.NewReplacer(
		"{{title}}", args.Title,
		"{{location}}", args.Location,
		"{{link}}", args.Link,
		"{{publishedOn}}", args.Published,
		"{{type}}", args.Type,
		"{{content}}", args.Content,
	)

	return replacer.Replace(templateContent), nil
}
