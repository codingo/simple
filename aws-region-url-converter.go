package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	regions := getRegions()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		region := getRegionFromURL(url)
		if region != "" {
			for _, r := range regions {
				if r != region {
					newURL := strings.Replace(url, region, r, -1)
					fmt.Println(newURL)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}

func getRegions() []string {
	resp, err := http.Get("https://docs.aws.amazon.com/general/latest/gr/rande.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get AWS regions:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	regions := []string{}
	re := regexp.MustCompile(`>([a-z]{2}-[a-z]+-\d)<`)
	buf := make([]byte, 4*1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "failed to read response body:", err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}

		matches := re.FindAllSubmatch(buf[:n], -1)
		for _, match := range matches {
			regions = append(regions, string(match[1]))
		}
	}

	return regions
}

func getRegionFromURL(url string) string {
	re := regexp.MustCompile(`[a-z]{2}-[a-z]+-\d`)
	match := re.FindString(url)
	return match
}
