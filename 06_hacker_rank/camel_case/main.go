package main

import (
	"fmt"
)

func main() {

	var st string
	fmt.Scanf("%s\n", &st)

	numberWords := countCamelCaseWords(st)
	fmt.Println(numberWords)
}

func countCamelCaseWords(st string) int {
	count := 0
	for i, ch := range st {
		// Check if the char is uppercase or if the first char is lowercase
		if (ch >= 'A' && ch <= 'Z') || (i == 0 && (ch >= 'a' && ch <= 'z')) {
			count++
		}
	}
	return count
}
