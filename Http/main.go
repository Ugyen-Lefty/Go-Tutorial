package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("error ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Number of bytes = ", bs)
	return len(bs), nil
}
