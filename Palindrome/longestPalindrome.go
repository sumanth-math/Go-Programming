// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	//AbhjhKLKNMNMNnmbmbmbm

	input := "AbhjhKLKNMNMNnmbmbmbm"

	data := strings.ToLower(input)

	longestPalindrome := ""
	for i := 0; i < len(data); i++ {
		var sb1 strings.Builder
		sb1.WriteString(string(data[i]))
		for j := i + 1; j < len(data); j++ {
			sb1.WriteString(string(data[j]))
			subString := sb1.String()
			isPalindrome := checkPalindrome(subString)
			if isPalindrome {
				if len(longestPalindrome) < len(subString) {
					longestPalindrome = subString
				}
			}
		}
	}

	fmt.Printf("Longest Palindrome in %s is %s\n", input, longestPalindrome)

}

func checkPalindrome(str string) bool {

	if len(str) == 0 ||  len(str) == 1 {
		return false
	}

	palindrome := true
	middle := 0
	if len(str)%2 == 0 {
		middle = len(str) / 2
	} else {
		middle = len(str)/2 + 1
	}

	for i := 0; i < middle; i++ {
		if str[i] != str[len(str)-1-i] {
			palindrome = false
			break
		}
	}

	return palindrome
}
