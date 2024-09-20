package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sortCharacters(s string) (string, string) {
	vowels := "aiueo"
	var vowelsCars, consonanChart string

	s = strings.ToLower(s)
	for _, char := range s {
		if char == ' ' {
			continue
		}
		if strings.ContainsRune(vowels, char) {
			vowelsCars += string(char)
		} else if char >= 'a' && char <= 'z' {
			consonanChart += string(char)
		}
	}

	return vowelsCars, consonanChart
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	vowels, consonants := sortCharacters(input)
	fmt.Println("Vowel Characters:", vowels)
	fmt.Println("Consonant Characters:", consonants)
}
