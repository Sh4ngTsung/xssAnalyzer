package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	searchString := flag.String("p", "", "")
	flag.Parse()

	if *searchString == "" {
		fmt.Println("Error: Please provide a search string using the -p flag")
		fmt.Println("Usage: xssAnalyzer -p \"<svg onload=confirm(1)>\"")
		return
	}

	x := bufio.NewScanner(os.Stdin)

	jobs := make(chan string)
	var wg sync.WaitGroup

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for domain := range jobs {
				func() {
					resp, err := client.Get(domain)
					if err != nil {
						return
					}
					defer resp.Body.Close()

					body, err := io.ReadAll(resp.Body)
					if err != nil {
						fmt.Println(err)
						return
					}
					sb := string(body)

					if strings.Contains(sb, *searchString) {
						fmt.Println(colorRed, "XSS FOUND:", domain, colorReset)
					} else {
						fmt.Println(colorGreen, "NOT VULN:", domain, colorReset)
					}
				}()
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
