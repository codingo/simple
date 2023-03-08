package main

import (
    "bufio"
    "fmt"
    "io"
    "net/url"
    "os"
    "strings"
)

var regionShortCodes = []string{
    "us-east-1", "us-east-2", "us-west-1", "us-west-2",
    "af-south-1", "ap-east-1", "ap-south-1", "ap-northeast-3", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2",
    "ca-central-1", "eu-central-1", "eu-west-1", "eu-west-2", "eu-south-1", "eu-west-3",
    "eu-north-1", "me-south-1", "sa-east-1",
}

func main() {
    // create a scanner to read from standard input
    scanner := bufio.NewScanner(os.Stdin)

    // read each line from standard input
    for scanner.Scan() {
        line := scanner.Text()

        // parse the URL to get the hostname and the path
        u, err := url.Parse(line)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing URL %q: %v\n", line, err)
            continue
        }

        // split the hostname into the domain and the region short code
        parts := strings.SplitN(u.Hostname(), "-", 2)
        if len(parts) != 2 {
            fmt.Fprintf(os.Stderr, "Error parsing region short code from URL %q\n", line)
            continue
        }

        domain := parts[0]
        regionShortCode := parts[1]

        // iterate over all the region short codes and output a copy of the URL for each one
        for _, code := range regionShortCodes {
            if code != regionShortCode {
                newHost := fmt.Sprintf("%s-%s.%s", domain, code, u.Port())
                newURL := url.URL{
                    Scheme:   u.Scheme,
                    Host:     newHost,
                    Path:     u.Path,
                    RawQuery: u.RawQuery,
                    Fragment: u.Fragment,
                }
                fmt.Println(newURL.String())
            }
        }
    }

    if err := scanner.Err(); err != nil {
        if err == io.EOF {
            return
        }
        fmt.Fprintf(os.Stderr, "Error reading from standard input: %v\n", err)
    }
}
