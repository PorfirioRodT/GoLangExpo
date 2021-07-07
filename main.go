package main

import (
	"sort"
	"strings"
)

func sortingValues(word string) string {

	words := strings.Split(word, "")
	sort.Strings(words)
	return strings.Join(words, "")

}

func main() {

}
