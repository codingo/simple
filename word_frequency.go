package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type WordFrequency struct {
	word  string
	count int
}

func main() {
	targets := flag.String("t", "", "targets file (newline per webpage to load)")
	exclude := flag.String("e", "", "exclude file (newline per word to exclude)")
	number := flag.Int("n", 10, "the number of most common words to output")
	threads := flag.Int("threads", 10, "the number of threads to use")

	flag.Parse()

	if *targets == "" {
		fmt.Println("Error: Missing -t or -targets flag")
		os.Exit(1)
	}

	excludedWords, err := loadExcludedWords(*exclude)
	if err != nil {
		fmt.Println("Error loading excluded words:", err)
		os.Exit(1)
	}

	urls, err := loadURLs(*targets)
	if err != nil {
		fmt.Println("Error loading URLs:", err)
		os.Exit(1)
	}

	urlsChan := make(chan string)
	go func() {
		for _, url := range urls {
			urlsChan <- url
		}
		close(urlsChan)
	}()

	var wg sync.WaitGroup
	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urlsChan {
				wordList, err := extractWords(url)
				if err != nil {
					fmt.Println("Error extracting words:", err)
					continue
				}

				wordMap := make(map[string]int)
				for _, word := range wordList {
					word = strings.ToLower(word)
					if !excludedWords[word] {
						wordMap[word]++
					}
				}

				frequencies := createFrequencies(wordMap)

				fmt.Printf("\nResults for %s:\n", url)
				printFrequencies(frequencies, *number)
			}
		}()
	}

	wg.Wait()
}
