package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortCharacters(s string) (string, string) {
	vowels := "aiueo"
	var vowelsCars, consonanChart string

	s = strings.ToLower(s)
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			vowelsCars += string(char)
		} else if char >= 'a' && char <= 'z' {
			consonanChart += string(char)
		}
	}

	vowelsSlice := strings.Split(vowelsCars, "")
	consonantsSlice := strings.Split(consonanChart, "")
	sort.Strings(vowelsSlice)
	sort.Strings(consonantsSlice)

	return strings.Join(vowelsSlice, ""), strings.Join(consonantsSlice, "")
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
