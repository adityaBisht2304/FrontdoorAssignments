package course1week2

import (
	"fmt"
	"strings"
)

func Question_1() {
	var numFloat float32
	fmt.Printf("Enter a floating point number")
	fmt.Scan(&numFloat)

	var numTruncated int = int(numFloat)
	fmt.Printf("Truncated version of floating point number is: %v", numTruncated)
}

func searchSpecificCharacters(str string) bool {
	start := 0
	end := len(str) - 1

	if str[start] != 'i' {
		return false
	}

	if str[end] != 'n' {
		return false
	}

	for i := 1; i < end; i++ {
		if str[i] == 'a' {
			return true
		}
	}

	return false
}

func Question_2() {
	var str string
	fmt.Printf("Enter a string")
	fmt.Scan(&str)

	if searchSpecificCharacters(strings.ToLower(str)) {
		fmt.Printf("Found")
	} else {
		fmt.Printf("Not Found")
	}
}
