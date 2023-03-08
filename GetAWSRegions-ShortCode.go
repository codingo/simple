package main

import (
    "fmt"
    "net/http"
    "strings"
    "golang.org/x/net/html"
)

func main() {
    // send a GET request to the Amazon RDS documentation page
    resp, err := http.Get("https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // parse the HTML content of the page using the Golang HTML package
    doc, err := html.Parse(resp.Body)
    if err != nil {
        panic(err)
    }

    // find the section of the page that contains the list of regions
    regionSection := findRegionSection(doc)

    // extract the list of regions from the section
    regions := extractRegions(regionSection)

    // print the list of regions to the console
    fmt.Println(regions)
}

func findRegionSection(n *html.Node) *html.Node {
    // find the div element with the id "concept-regions"
    if n.Type == html.ElementNode && n.Data == "div" {
        for _, attr := range n.Attr {
            if attr.Key == "id" && attr.Val == "concept-regions" {
                return n
            }
        }
    }

    // recursively search for the region section in the child nodes
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        result := findRegionSection(c)
        if result != nil {
            return result
        }
    }

    return nil
}

func extractRegions(regionSection *html.Node) []string {
    var regions []string

    // find all the links in the region section and extract the region ID from the href attribute
    for c := regionSection.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.ElementNode && c.Data == "a" {
            regionName := strings.TrimSpace(c.FirstChild.Data)
            if regionName != "China (Beijing)" && regionName != "China (Ningxia)" {
                // exclude China regions as they require separate signup
                href := c.Attr[0].Val
                regionID := strings.Split(href, "/")[5]
                regions = append(regions, regionID)
            }
        }
    }

    return regions
}
