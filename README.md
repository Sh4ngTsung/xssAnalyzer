#Description

The program above is a simple XSS (Cross-Site Scripting) vulnerability analyzer that checks whether a list of domains is vulnerable to XSS attacks.

The program reads a list of domains from the standard input and then starts multiple goroutines to check if each domain is vulnerable to XSS. To do this, it sends an HTTP request to each domain and checks if the response contains one or more strings that are common in XSS attacks.

If the program finds an XSS string in the server's response, it will print "XSS FOUND:" followed by the corresponding domain in red. Otherwise, it will print "Not Vulnerable:" followed by the corresponding domain in green. The program uses different colors to visually highlight the results.

The program allows you to specify the string to be searched in the HTTP response using the "-p" flag. This allows you to search for different XSS patterns and adjust the program to look for specific vulnerabilities in a particular domain.

##Installation

To install the program, you need to have Go installed on your machine.

go install github.com/sh4ngtsung/xssanalyzer@latest

##Usage

To use the program, you need to provide a list of domains to analyze. The list should be provided via the standard input.

For example, to analyze a list of domains stored in a file called "domains.txt", you can run the following command:

cat domains.txt | waybackurls | gf xss | qsreplace '\<img src=x onerror=confirm(1)>' | xssAnalyzer -p "confirm(1)"
