package main

import (
	"fmt"
)

const (
	ASCII_UPPER_TO_LOWER_DIFF = 32
)

func main() {
	// Get input from the user
	var input string
	fmt.Print("Enter a string to check if it's a palindrome: ")
	fmt.Scanln(&input)

	// Call the IsPalindrome function and print the result
	result := IsPalindrome(input)
	if result {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}

// IsPalindrome checks if a string is a palindrome.
// A palindrome reads the same backward as forward, ignoring case, spaces, and punctuation.
func IsPalindrome(s string) bool {
	palindrome := ""
	for _, word := range extractWords(s) {
		palindrome += word
	}

	if len(palindrome) == 0 {
		return true
	}

	// We only need to check one half, as we're cheking both ends at the same
	// time. Should the length of the word be be odd, the middle character will
	// be considered "shared" so it can be ignored
	for idx := 0; idx <= (len(palindrome)-1)/2; idx++ {
		if palindrome[idx] != palindrome[len(palindrome)-1-idx] {
			return false
		}
	}

	return true
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
