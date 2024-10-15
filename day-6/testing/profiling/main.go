package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AnalyzeText takes a filename as input and reads the content of the file.
// It loops over each word in the `words` slice and performs the following actions:
// - Increments the count for the word in the `frequency` map.
// - Updates the `longest` word if the current word is longer.
// - Adds the length of the current word to `totalLength`.
// After the loop, it calculates the average word length by dividing `totalLength` by the length of the `words` slice.
// Finally, it returns a TextAnalysis struct with the following properties assigned:
// - Content: The original text content of the file.
// - WordFrequency: A map containing the frequency of occurrence of each word.
// - AvgWordLength: The average length of words in the text.
// - LongestWord: The longest word found in the text.

type TextAnalysis struct {
	Content       string
	WordFrequency map[string]int
	AvgWordLength float64
	LongestWord   string
}

func AnalyzeText(filename string) (TextAnalysis, error) {
	// reading the entire file in memory
	content, err := os.ReadFile(filename)
	if err != nil {
		return TextAnalysis{}, err
	}

	// Convert the byte slice of the file content to a string.
	text := string(content)

	// Split the `text` string into a slice of words.
	words := strings.Fields(text)

	// Initialize a map to hold the frequency of occurrence of each word.
	frequency := make(map[string]int)

	// Initialize totalLength to cumulate the length of all words.
	totalLength := 0

	// Initialize longest to hold the longest word.
	longest := ""

	// Loop over every word in the `words` slice.
	for _, word := range words {
		frequency[word]++ // Increment the count for the word in the frequency map.
		if len(word) > len(longest) {
			longest = word // Update the longest word if the current word is longer.
		}
		totalLength += len(word) // Add the length of the current word to totalLength.
	}

	// Calculating the average word length.
	avgLength := float64(totalLength) / float64(len(words))

	// Return a TextAnalysis instance with properties assigned.
	return TextAnalysis{
		Content:       text,
		WordFrequency: frequency,
		AvgWordLength: avgLength,
		LongestWord:   longest,
	}, nil

}

var cache = make(map[string]TextAnalysis)

//even if you see '0 allocations' during the benchmark process, allocations are still happening, but only on the first run. On subsequent runs for the same file, the allocations aren't happening not because our code got optimally efficient, but because we are reusing previously calculated results thanks to caching.
//Take into account that cache memory is also a part of the overall memory and contributes to memory usage. If you are processing large files or several files, your memory usage will increase significantly due to the populated cache.
//In conclusion, caching will not reduce the number of allocations for each unique input (i.e., for each new file that you analyze). It will only reduce allocations for repeat inputs (i.e., when you analyze an already-analyzed file again) by reusing the result from the previous run instead of re-calculating it.

func OptimizedAnalyzeText(filename string) (TextAnalysis, error) {
	if res, ok := cache[filename]; ok {
		return res, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return TextAnalysis{}, err
	}
	defer file.Close()

	frequency := make(map[string]int, 1000000)
	totalLength := 0
	longest := ""

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		frequency[word]++
		if len(word) > len(longest) {
			longest = word
		}
		totalLength += len(word)
	}

	if scanner.Err() != nil {
		return TextAnalysis{}, scanner.Err()
	}

	avgLength := float64(totalLength) / float64(len(frequency))
	t := TextAnalysis{
		WordFrequency: frequency,
		AvgWordLength: avgLength,
		LongestWord:   longest,
	}
	cache[filename] = t
	return t, nil

}

func main() {
	filename := "moby.txt"
	analysis, err := AnalyzeText(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	//fmt.Println("Word Frequency:", analysis.WordFrequency)
	fmt.Println("Average Word Length:", analysis.AvgWordLength)
	fmt.Println("Longest Word:", analysis.LongestWord)

}
