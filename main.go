/* Print http response codes of multiple url's from file */

package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("Starting multiurl-checker")

	content, err := os.ReadFile("urllist")

	if err != nil {
		log.Fatal("urllist file not found! it's required")
	}

	urls := strings.Split(string(content), "\n")

	log.Println("Check starting, going to check", len(urls), "urls")

	for _, element := range urls {
		// program expects urls without protocol. this might be improved with url.Parse
		resp, err := http.Get("http://" + element)

		if err != nil {
			log.Println("err: ", err)
		}

		if resp != nil {
			log.Println(resp.Request.URL, "\tstatus\t", resp.StatusCode)
		}
	}
}
