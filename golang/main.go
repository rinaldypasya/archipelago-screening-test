package main

import (
	"fmt"
	"regexp"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Remove punctuation using regex
	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")

	// Split text into words
	words := strings.Fields(text)

	// Create a map to store word frequencies
	frequency := make(map[string]int)

	// Count word occurrences
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	text := "Four, One two two three Three three four  four   four"
	freq := wordFrequency(text)

	// Print word frequencies
	for word, count := range freq {
		fmt.Printf("%s => %d\n", word, count)
	}
}
