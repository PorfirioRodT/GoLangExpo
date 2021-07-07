package main

import (
	"hash/fnv"
	"sort"
	"strings"
)

func sortingValues(word string) string {

	words := strings.Split(word, "")
	sort.Strings(words)
	return strings.Join(words, "")

}

func hashValues(values string) int32 {

	sortedValues := sortingValues(values)

	hash := fnv.New32()
	hash.Write([]byte(sortedValues))

	data := hash.Sum32()

	return int32(data)

}

func main() {

}
