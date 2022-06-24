package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.baidu.com",
		"https://pkg.go.dev",
		"http://stackoverflow.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, ch)
		}(l)
		// go checkLink(l, ch)
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("ping error:", link)
		ch <- link
		return
	}
	fmt.Println("ping success", link)
	ch <- link
}
