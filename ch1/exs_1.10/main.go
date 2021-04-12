package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type response struct {
	err  error
	data string
}

func (r response) String() string {
	if r.err != nil {
		return fmt.Sprint(r.err)
	} else {
		return r.data
	}
}

func main() {
	headers, urls := parseArgs(os.Args[1:])
	headerKeyValues := splitHeaders(headers)
	start := time.Now()
	ch := make(chan response)
	for _, url := range urls {
		go fetch(url, headerKeyValues, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, headers map[string]string, ch chan<- response) {
	start := time.Now()
	resp, err := get(url, headers)
	if err != nil {
		ch <- response{err: err}
		return
	}
	buffer := new(bytes.Buffer)
	nbytes, err := io.Copy(buffer, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- response{err: fmt.Errorf("fetch %s: %w", url, err)}
		return
	}
	secs := time.Since(start).Seconds()
	ch <- response{data: fmt.Sprintf("* * *\n"+
		"%s\n"+
		"%.2fs  %7d  %s\n"+
		"- - -",
		buffer, secs, nbytes, url,
	)}
}

func parseArgs(args []string) (headers []string, urls []string) {
	parseHeader := false
	for _, arg := range args {
		if strings.ToUpper(arg) == "-H" {
			parseHeader = true
		} else if parseHeader {
			headers = append(headers, arg)
			parseHeader = false
		} else {
			urls = append(urls, arg)
		}
	}
	return
}

func get(url string, headers map[string]string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	return http.DefaultClient.Do(req)
}

func splitHeaders(headers []string) map[string]string {
	keyValues := make(map[string]string)
	for _, header := range headers {
		if keyValue := strings.Split(header, ":"); len(keyValue) == 2 {
			keyValues[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
		}
	}
	return keyValues
}
