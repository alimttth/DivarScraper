

# Divar Scraper

## Introduction

Divar Scraper is a GoLang application that periodically scrapes the Divar website to find posts related to a specific keyword (in this case, MacBook) within a certain price range. When a matching post is found, it creates a system alert with the post's details.

## Features

- Periodically fetches data from the Divar website
- Filters posts based on a target keyword and price range
- Generates system alerts for matching posts

## Installation

To use Divar Scraper, follow these steps:

1. Make sure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).

2. Clone the repository:

```bash
git clone https://github.com/your_username/divar-scraper.git

```
## Running

```bash

cd divar-scraper

go mod tidy

go build

./divar-scraper

```


# Configuration

You can configure the scraper by modifying the following variables in the main function of main.go:

-BASE_URL: The base URL of the Divar website with search parameters.
-TARGET_TITLE: The target keyword to search for in post titles.


#create Readme ChatGPT
