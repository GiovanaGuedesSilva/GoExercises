package main

/*
	Create a map that counts how many times each word appears in a sentence.
*/

import (
	"fmt"
	"strings"
)

func main() {
	// Define the sentence
	fmt.Println("Sentence: [go é ótimo e go é divertido]")
	sentence := "go é ótimo e go é divertido"

	// Split the sentence into words
	words := strings.Fields(sentence)

	// Create a map to count the frequency of each word
	wordCount := make(map[string]int)

	// Iterate over the words and update the count in the map
	for _, word := range words {
		wordCount[word]++
	}

	// Print the word frequencies
	fmt.Println("\nWord count:\n")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
