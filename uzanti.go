package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Global değişkenler
var (
	baseURL     string
	directories []string
	extensions  []string
)

var MAIN string = "\x1b[38;5;85m"
var GREEN string = "\x1b[38;5;82m"
var GRAY string = "\x1b[38;5;246m"
var NAME string = "\x1b[38;5;228m"
var RED string = "\x1b[1;31m"
var FAIL string = "\x1b[1;91m"
var ORANGE string = "\x1b[0;38;5;214m"
var LRED string = "\x1b[0;38;5;196m"
var BOLD string = "\x1b[1m"
var PURPLE string = "\x1b[0;38;5;141m"
var BLUE string = "\x1b[0;38;5;12m"
var UNDERLINE string = "\x1b[4m"
var UNSTABLE string = "\x1b[5m"
var END string = "\x1b[0m"

func ikinciasama(directory []string) {
	payload := "%s*~1*.%s"
	for _, dir := range directory {
		for _, ext := range extensions {
			target := baseURL + fmt.Sprintf(payload, dir, ext) 
			req, err := http.NewRequest("OPTIONS", target, nil)
			if err != nil {
				fmt.Println
				continue
			}

			// Send the request using the proxy settings
			proxyURL, _ := url.Parse("http://127.0.0.1:8080")
			transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
			client := &http.Client{Transport: transport}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == 404 {
				last := fmt.Sprintf("%s%sURL\t\t\t\t\t\tDIRECTORY/FILE\t\t\t\tEXT%s\n", BOLD, ORANGE, END)
				fmt.Println(last)
				output := fmt.Sprintf("%s%s\t\t%s\t\t\t\t\t%s%s", MAIN, target, dir, ext, END)
				fmt.Println(output)
			}
		}
	}
}

func main() {
	flag.StringVar(&baseURL, "url", "", "URL'yi belirtin") 
	directoriesFlag := flag.String("dir", "", "Virgülle ayrilmis dizin listesi")
	extensionsFlag := flag.String("ext", "txt,jpg,png,htm", "Virgülle ayrilmis dosya uzantilari listesi")
	flag.Parse()

	// Virgülle ayrılmış string'i diziye çevir
	extensions = strings.Split(*extensionsFlag, ",")
	if *directoriesFlag != "" {
		directories = strings.Split(*directoriesFlag, ",")
	}

	// Komut satırı argümanlarından URL'yi al
	if len(flag.Args()) > 0 {
		baseURL = flag.Arg(0)
	}

	if baseURL == "" {
		fmt.Println("Lütfen -url bayragi ile bir URL belirtin")
		os.Exit(1)
	}

	ikinciasama(directories)
}
