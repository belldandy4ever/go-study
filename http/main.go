package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWrite struct{}

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lw := logWrite{}

	io.Copy(lw, resp.Body)

}

func (logWrite) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println(len(string(bs)))
	return len(bs), nil
}
