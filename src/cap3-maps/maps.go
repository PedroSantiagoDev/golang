package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]

	statistics := collectStatistics(words)

	printStatistics(statistics)
}

func collectStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		accountant, found := statistics[initial]

		if found {
			statistics[initial] = accountant + 1
		} else {
			statistics[initial] = 1
		}
	}
	return statistics
}

func printStatistics(statistics map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letra:")

	for initial, counter := range statistics {
		fmt.Printf("%s = %d\n", initial, counter)
	}
}
