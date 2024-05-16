package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			break
		}
		fmt.Println("Input Text: ", input)
	}
}
