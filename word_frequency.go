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

	wordMap := make(map[string]int)

	for _, url := range urls {
		wordList, err := extractWords(url)
		if err != nil {
			fmt.Println("Error extracting words:", err)
			continue
		}

		for _, word := range wordList {
			word = strings.ToLower(word)
			if !excludedWords[word] {
				wordMap[word]++
			}
		}
	}

	frequencies := createFrequencies(wordMap)

	printFrequencies(frequencies, *number)
}

func loadExcludedWords(filename string) (map[string]bool, error) {
	excludedWords := make(map[string]bool)

	if filename == "" {
		return excludedWords, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		excludedWords[strings.ToLower(scanner.Text())] = true
	}

	return excludedWords, nil
}

func loadURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	urls := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	return urls, nil
}

func createFrequencies(wordMap map[string]int) []WordFrequency {
	frequencies := []WordFrequency{}

	for word, count := range wordMap {
		frequencies = append(frequencies, WordFrequency{word, count})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].count > frequencies[j].count
	})

	return frequencies
}

func printFrequencies(frequencies []WordFrequency, number int) {
	for i := 0; i < number && i < len(frequencies); i++ {
		fmt.Printf("%s: %d\n", frequencies[i].word, frequencies[i].count)
	}
}

func extractWords(url string) ([]string, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	words := []string{}
	doc.Find("body").Each(func(_ int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		words = append(words, strings.Split(text, " ")...)
	})

	return words, nil
}
