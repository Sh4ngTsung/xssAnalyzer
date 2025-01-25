# XSS Analyzer

The **XSS Analyzer** is a simple **Cross-Site Scripting (XSS)** vulnerability scanner. It checks if a list of domains is vulnerable to XSS attacks by sending HTTP requests and looking for typical XSS patterns in the responses.

## Description

This program reads a list of domains from the standard input and then uses **goroutines** to check if each domain is vulnerable to XSS. It does this by sending HTTP requests to each domain and checking if the response contains one or more strings commonly associated with XSS attacks.

- **XSS FOUND**: If an XSS pattern is found, the program prints `XSS FOUND:` followed by the domain in red.
- **Not Vulnerable**: Otherwise, it prints `Not Vulnerable:` followed by the domain in green.

You can customize the XSS patterns to search for in the HTTP response using the `-p` flag, allowing you to tailor the scan to specific vulnerabilities in a given domain.

## Installation

### Using Go

To install the program using Go, make sure Go is installed on your machine.

Run the following command to install:

```
go install github.com/sh4ngtsung/xssanalyzer@latest
```

### Using Git Clone
```
git clone https://github.com/Sh4ngTsung/xssAnalyzer.git
cd xssAnalyzer
go build -ldflags "-s -w" -o xssAnalyzer xssAnalyzer.go
./xssAnalyzer
```

### Usage


To analyze a list of domains stored in a file called domains.txt, run the following command:
```
cat domains.txt | waybackurls | gf xss | qsreplace '\<img src=x onerror=confirm(1)>' | xssAnalyzer -p "confirm(1)"
```

### Example 2: Using subdomain enumeration tools

You can also use tools like assetfinder, gauplus, and gf to collect domains/subdomains and analyze them with the XSS Analyzer. Example:
```
echo "example.com" | assetfinder -subs-only | gauplus | gf xss | qsreplace '\<svg onload=prompt(document.domain)>' | xssAnalyzer -p "prompt(document.domain)"
```
