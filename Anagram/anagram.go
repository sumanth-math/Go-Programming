package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	fmt.Println(Yellow + "**** Anagram Checker ****" + Reset)
	fmt.Println(Yellow + "Provide the first anangram text:" + Reset)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	anagramText1 := scanner.Text()

	fmt.Println(Yellow + "Provide the second anangram text:" + Reset)
	scanner.Scan()
	anagramText2 := scanner.Text()

	isAnagram := anagramCheck(anagramText1, anagramText2)
	if isAnagram {
		fmt.Printf(Green + fmt.Sprintf("\"%s\" and \"%s\" are Anagrams\n", anagramText1, anagramText2) + Reset)
	} else {
		fmt.Printf(Red + fmt.Sprintf("\"%s\" and \"%s\" are NOT Anagrams\n", anagramText1, anagramText2) + Reset)
	}

}

func anagramCheck(anagramText1 string, anagramText2 string) bool {
	prunedAnagramText1 := strings.Replace(strings.ToLower(anagramText1), " ", "", -1)
	prunedAnagramText2 := strings.Replace(strings.ToLower(anagramText2), " ", "", -1)

	if len(prunedAnagramText1) != len(prunedAnagramText2) {
		return false
	}

	dictionary := make(map[rune]int)
	for index := 0; index < len(prunedAnagramText1); index++ {
		if _, ok := dictionary[rune(prunedAnagramText1[index])]; ok {
			delete(dictionary, rune(prunedAnagramText1[index]))
		} else {
			dictionary[rune(prunedAnagramText1[index])] = 1
		}

		if _, ok := dictionary[rune(prunedAnagramText2[index])]; ok {
			delete(dictionary, rune(prunedAnagramText2[index]))
		} else {
			dictionary[rune(prunedAnagramText2[index])] = 1
		}
	}
	if len(dictionary) > 0 {
		return false
	}

	return true
}
