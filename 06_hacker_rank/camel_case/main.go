package main

import (
	"fmt"
)

func main() {
	st := "holaMundoDelCamelCase"
	numberWords := countCamelCaseWords(st)
	if numberWords == -1 {
		fmt.Println("The string provided was empty.")
	} else {
		fmt.Println(numberWords + 1)
	}
}

func countCamelCaseWords(st string) int {
	if len(st) == 0 {
		return -1
	}
	count := 0
	for _, ch := range st {
		if ch >= 65 && ch <= 90 {
			count++
		}
	}
	return count
}
