// Package challenge6 contains the solution for Challenge 6.
package challenge6

// Add any necessary imports here

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}

const (
	ASCII_UPPER_TO_LOWER_DIFF = 32
)

func CountWordFrequency(text string) map[string]int {
	freq := map[string]int{}

	for _, word := range extractWords(text) {
		if word != "" {
			freq[word]++
		}
	}

	return freq
}

func extractWords(text string) []string {
	words := []string{}
	wordCandidate := ""

	for _, char := range text {
		switch {
		case isAlphaNumeric(char):
			wordCandidate += string(toLower(char))
		case isSeparator(char) && len(wordCandidate) != 0:
			words = append(words, wordCandidate)
			wordCandidate = ""
		}
	}
	if len(wordCandidate) != 0 {
		words = append(words, wordCandidate)
	}

	return words
}

func isSeparator(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '-'
}

func isAlphaNumeric(char rune) bool {
	return (('A' <= char && char <= 'Z') ||
		('a' <= char && char <= 'z') ||
		('0' <= char && char <= '9'))
}

func toLower(char rune) rune {
	if 'A' <= char && char <= 'Z' {
		return rune(char + ASCII_UPPER_TO_LOWER_DIFF)
	}

	return char
}
