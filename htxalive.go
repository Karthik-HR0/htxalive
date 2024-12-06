package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	currentVersion = "v1.0.1"
	githubRepo     = "Karthik-HR0/htxalive"
)

var (
	output      string
	silent      bool
	concurrency int
	threads     int
	showStatus  bool
	showTitle   bool
	showIP      bool
)

func init() {
	flag.StringVar(&output, "o", "", "Collect only output URLs")
	flag.BoolVar(&silent, "s", false, "Silent mode")
	flag.IntVar(&concurrency, "c", 10, "Concurrency level")
	flag.IntVar(&threads, "t", 3, "Threading level")
	flag.BoolVar(&showStatus, "sc", false, "Show HTTP status codes")
	flag.BoolVar(&showTitle, "tl", false, "Show page titles")
	flag.BoolVar(&showIP, "ip", false, "Show IP address")
	flag.Parse()
}

func printLogo() {
	logo := `

.__     __ ____  ___      .__  .__              
|  |___/  |\   \/  /____  |  | |__|__  __ ____  
|  |  \   __\     /\__  \ |  | |  \  \/ // __ \ 
|   Y  \  | /     \ / __ \|  |_|  |\   /\  ___/ 
|___|  /__|/___/\  (____  /____/__| \_/  \___  >
     \/          \_/    \/                   \/ 
     
                           with <3 by @Karthik-HR0

`
	fmt.Println(logo)
}

func checkLatestVersion() string {
	url := fmt.Sprintf("https://github.com/Karthik-HR0/htxalive/releases/latest", githubRepo)
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// Extract version from JSON response
	re := regexp.MustCompile(`"tag_name":"([^"]+)"`)
	matches := re.FindStringSubmatch(string(body))
	
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func printVersionInfo() {
	latestVersion := checkLatestVersion()
	
	if latestVersion == "" {
		fmt.Printf("[INF] Current htxalive version %s (version check failed)\n", 
		                    
		currentVersion)
		return
	}

	if latestVersion == currentVersion {
		fmt.Printf("[INF] Current htxalive version %s \033[32m(latest)\033[0m\n", currentVersion)
	} else {
		fmt.Printf("[INF] Current htxalive version %s \033[31m(outdated)\033[0m\n", currentVersion)
		fmt.Printf("[INF] Latest version: %s\n", latestVersion)
	}
}

func fetchTitle(body []byte) string {
	re := regexp.MustCompile(`<title>(.*?)</title>`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) > 1 {
		return matches[1]
	}
	return " "
}

func resolveIP(url string) string {
	// Remove protocol for hostname extraction
	host := strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://")
	
	// Split host:port if present
	hostParts := strings.Split(host, ":")
	host = hostParts[0]

	ips, err := net.LookupIP(host)
	if err != nil || len(ips) == 0 {
		return ""
	}

	// Return first IPv4 address
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}

func httpxer(url string, wg *sync.WaitGroup, outputChan chan<- string) {
	defer wg.Done()
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 0 {
		var title string
		body, _ := ioutil.ReadAll(resp.Body)

		if showTitle {
			title = fetchTitle(body)
		}

		if output != "" {
			outputChan <- url
			return
		}

		// Resolve IP if -ip flag is set
		ipStr := ""
		if showIP {
			ipStr = resolveIP(url)
		}

		// Formatting with optional IP
		var outputStr string
		if showIP && ipStr != "" {
			outputStr = fmt.Sprintf("%s [\033[33m%s\033[0m]", url, ipStr)
		} else {
			outputStr = url
		}

		if showStatus {
			var color string
			switch {
			case resp.StatusCode >= 200 && resp.StatusCode < 300:
				color = "\033[32m"
			case resp.StatusCode >= 300 && resp.StatusCode < 400:
				color = "\033[33m"
			case resp.StatusCode >= 400 && resp.StatusCode < 600:
				color = "\033[31m"
			default:
				color = "\033[0m"
			}
			
			if title != "" {
				outputStr = fmt.Sprintf("%s [%s%d\033[0m] [\033[34;1m%s\033[0m]", 
					outputStr, color, resp.StatusCode, title)
			} else {
				outputStr = fmt.Sprintf("%s [%s%d\033[0m]", outputStr, color, resp.StatusCode)
			}
		} else if title != "" {
			outputStr = fmt.Sprintf("%s [\033[34;1m%s\033[0m]", outputStr, title)
		}

		fmt.Println(outputStr)
	}
}

func main() {
	// Print logo and version info
	printLogo()
	printVersionInfo()

	var storage []string

	// Read input from stdin or file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
			storage = append(storage, url)
		} else {
			storage = append(storage, "https://"+url, "http://"+url)
		}
	}

	// Output channel for -o flag
	outputChan := make(chan string, len(storage))

	var wg sync.WaitGroup
	for _, url := range storage {
		wg.Add(1)
		go httpxer(url, &wg, outputChan)
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	// Handle output file if -o flag is used
	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			return
		}
		defer file.Close()

		for url := range outputChan {
			fmt.Fprintln(file, url)
		}
	} else {
		// Consume output channel if not writing to file
		for range outputChan {}
	}
}
