package main

import (
	"hash/fnv"
	"log"
	"sort"
	"strings"
	"time"
	"log"
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

func timerOfProgram(startTimer time.Time, name string)  {

	timeDuration := time.Since(startTimer)
	log.Printf("%s the program took a time of: %s", name, timeDuration)
}

func main() {

	

}
