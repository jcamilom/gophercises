package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter camelCase string: ")
	st, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	st = strings.TrimSuffix(st, "\n")
	numberWords := countCamelCaseWords(st)
	if numberWords == -1 {
		fmt.Println("The string provided is empty.")
	} else {
		fmt.Println(numberWords)
	}
}

func countCamelCaseWords(st string) int {
	if len(st) == 0 {
		return -1
	}
	count := 0
	for i, ch := range st {
		// Check if the char is uppercase or if the first char is lowercase
		if (ch >= 65 && ch <= 90) || (i == 0 && (ch >= 97 && ch <= 122)) {
			count++
		}
	}
	return count
}
