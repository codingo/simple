package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "strings"
)

func main() {
    // Get the AWS regions from the documentation URL
    regions, err := getRegions()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error getting regions: %v\n", err)
        os.Exit(1)
    }

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        // Get the input URL and extract the region short code
        url := scanner.Text()
        shortCode := extractShortCode(url)

        // Generate a copy of the URL for every other region
        for _, region := range regions {
            if region != shortCode {
                newUrl := strings.Replace(url, shortCode, region, -1)
                fmt.Println(newUrl)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
        os.Exit(1)
    }
}

// getRegions fetches the list of AWS regions from the documentation URL
func getRegions() ([]string, error) {
    url := "https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html"
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching regions: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }

    // Find all the region short codes in the body
    var regionRegex = regexp.MustCompile(`([a-z]{2}-[a-z]+-\d+|[a-z]{2}-\w+-\d+)`)
    shortCodes := regionRegex.FindAllString(string(body), -1)

    return shortCodes, nil
}

// extractShortCode extracts the short code (e.g. "us-west-1") from a URL
func extractShortCode(url string) string {
    var shortCodeRegex = regexp.MustCompile(`([a-z]{2}-[a-z]+-\d+|[a-z]{2}-\w+-\d+)`)
    match := shortCodeRegex.FindStringSubmatch(url)
    if len(match) > 1 {
        return match[1]
    }
    return ""
}
