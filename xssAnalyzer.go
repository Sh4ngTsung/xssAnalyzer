package main

import (
        "bufio"
        "flag"
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "strings"
        "sync"
)

func main() {
        colorReset := "\033[0m"
        colorRed := "\033[31m"
        colorGreen := "\033[32m"

        // Define a flag "p" to hold the search string
        searchString := flag.String("p", "", "string to search in response body")
        flag.Parse()

        // Ensure that a search string was provided
        if *searchString == "" {
                fmt.Println("Error: Please provide a search string using the -p flag")
                fmt.Println("Usage: xssAnalyzer -p \"<svg onload=confirm(1)>\"")
                return
        }

        x := bufio.NewScanner(os.Stdin)

        jobs := make(chan string)
        var wg sync.WaitGroup

        for i := 0; i < 30; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        for domain := range jobs {
                                resp, err := http.Get(domain)
                                if err != nil {
                                        continue
                                }
                                body, err := ioutil.ReadAll(resp.Body)
                                if err != nil {
                                        fmt.Println(err)
                                }
                                sb := string(body)

                                if strings.Contains(sb, *searchString) {
                                        fmt.Println(string(colorRed), "XSS FOUND:", domain, string(colorReset))
                                } else {
                                        fmt.Println(string(colorGreen), "NOT VULN:", domain, string(colorReset))
                                }
                        }
                }()
        }

        for x.Scan() {
                domain := x.Text()
                jobs <- domain
        }

        close(jobs)
        wg.Wait()
}
