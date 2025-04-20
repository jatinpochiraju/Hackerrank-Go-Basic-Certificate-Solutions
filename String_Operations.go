package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ModifyString processes the input string.
func ModifyString(str string) string {
	str = strings.TrimSpace(str)

	var noDigits []rune
	for _, r := range str {
		if !unicode.IsDigit(r) {
			noDigits = append(noDigits, r)
		}
	}

	// Reverse the runes
	for i, j := 0, len(noDigits)-1; i < j; i, j = i+1, j-1 {
		noDigits[i], noDigits[j] = noDigits[j], noDigits[i]
	}

	return string(noDigits)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	result := ModifyString(input)
	fmt.Println(result)
}
