package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			recordLocations(lines(counts), file, locations)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\t%s\n", strings.Join(uniq(locations[line]), ", "), n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func recordLocations(lines []string, file string, locations map[string][]string) {
	for _, line := range lines {
		if locations[line] == nil {
			locations[line] = make([]string, 0)
		}
		locations[line] = append(locations[line], file)
	}
}

func lines(counts map[string]int) []string {
	keys := make([]string, len(counts))
	i := 0
	for key := range counts {
		keys[i] = key
		i++
	}
	return keys
}

func uniq(strings []string) []string {
	set := make(map[string]struct{})
	uniqStrings := make([]string, 0, len(strings))
	for _, str := range strings {
		if _, ok := set[str]; ok {
			continue
		}
		set[str] = struct{}{}
		uniqStrings = append(uniqStrings, str)
	}
	return uniqStrings
}
