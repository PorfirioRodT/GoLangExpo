package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"sort"
	"strings"
	"time"
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

func timerOfProgram(startTimer time.Time, name string) {

	timeDuration := time.Since(startTimer)
	log.Printf("%s the program took a time of: %s", name, timeDuration)
}

func main() {

	defer timerOfProgram(time.Now(), "anagrams")

	file, err := os.Open("wordList.txt")

	if err != nil {

		log.Fatal(err)

	}

	defer file.Close()

	wordsMaps := make(map[int]map[int32][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		words := scanner.Text()
		wordsLenght := len(words)
		hashedWords := hashValues(words)
		wordsList := wordsMaps[wordsLenght][hashedWords]
		wordsList = append(wordsList, words)

		if wordsMaps[wordsLenght] == nil {

			temporalHash := make(map[int32][]string)
			temporalHash[hashedWords] = []string{words}
			wordsMaps[wordsLenght] = temporalHash

		} else {

			wordsMaps[wordsLenght][hashedWords] = wordsList

		}
	}

	file, err = os.Open("wordList.txt")

	if err != nil {

		log.Fatal(err)

	}

	defer file.Close()

	scanner = bufio.NewScanner(file)

	setsOfAnagramsQuantity := 0

	for scanner.Scan() {

		words := scanner.Text()
		wordsLenght := len(words)
		hashedWords := hashValues(words)

		if len(wordsMaps[wordsLenght][hashedWords]) > 1 {

			fmt.Println(wordsMaps[wordsLenght][hashedWords])
			wordsMaps[wordsLenght][hashedWords] = nil

			setsOfAnagramsQuantity++

		}

	}

	fmt.Println("The quantity of the sets is: ", setsOfAnagramsQuantity)

}
