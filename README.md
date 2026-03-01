# asnr-report-scrapper

A CLI tool written in Go to scrape and download nuclear incident reports (avis d'incident) from the French Nuclear Safety Authority website ([ASNR](https://reglementation-controle.asnr.fr/controle/actualites-du-controle/installations-nucleaires/avis-d-incident-des-installations-nucleaires)).

> **Note:** This project is currently a work in progress and not yet stable.

## Features

- Auto-detect the available year range of reports from the ASNR website
- Filter reports by a custom year range (`--from` / `--to`)
- Download reports concurrently
- Convert downloaded HTML reports to Markdown files using a simple template engine

## Requirements

- Go 1.25.5 or later

## Installation

```bash
git clone https://github.com/thegostisdead/asnr-report-scrapper.git
cd asnr-report-scrapper
go build -o scrapper
```

## Usage

```
./scrapper --help
./scrapper --from=2000 --to=2010
./scrapper --from=2000
./scrapper --to=2010
./scrapper --from=2000 --to=2010 --out reports
```

### Flags

| Flag     | Description                                      | Default                   |
|----------|--------------------------------------------------|---------------------------|
| `--from` | Start year for the report range                  | Auto-detected from ASNR   |
| `--to`   | End year for the report range                    | Auto-detected from ASNR   |
| `--out`  | Output directory where Markdown files are saved  | Executable directory      |

If neither `--from` nor `--to` is provided, the tool automatically detects the available interval by querying the ASNR website.

## Output

Each downloaded report is rendered as a Markdown file with the following front matter:

```markdown
---
title: <report title>
location: <nuclear installation location>
link: <source URL>
publishedOn: <publication date>
type: <incident type>
---
<report content>
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -v
```

CI runs automatically on push and pull requests to `main` via GitHub Actions.

## License

This project does not currently include a license file. All rights reserved by the author.
