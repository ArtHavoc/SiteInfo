package site

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// <-chan - canal somente leitura

func GetSiteName(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		//fmt.Println(url)
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func PrintChannel(ch <-chan string) {
	for c := range ch {
		fmt.Println(c)
	}
}
