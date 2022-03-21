// Fetch выводит ответ на запрос по заданному URL
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		s := resp.Status
		resp.Body.Close()
		fmt.Printf("%s", s)
	}
}
