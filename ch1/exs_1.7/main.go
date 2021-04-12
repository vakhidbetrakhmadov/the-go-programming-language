package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		for {
			written, err := io.Copy(os.Stdout, resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s: %v\n", os.Args[0], url, err)
				os.Exit(1)
			}
			if written == 0 {
				break
			}
		}
	}
}
