package main

import (
	"fmt"
)

const a = 'a'
const z = 'z'
const aM = 'A'
const zM = 'Z'

func main() {
	var n, k int
	var s string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%d\n", &k)

	fmt.Println(cipherString(s, k))
}

func cipherString(st string, k int) string {
	// In case the shift is bigger than the alphabet width, "normalize" it so
	// that it's never bigger than the alphabet width
	k = k % 26
	// If the shift is zero, do nothing
	if k == 0 {
		return st
	}
	runes := []rune(st)
	for i, ch := range runes {
		// Check if the char is not a letter
		if min, may := (ch >= a && ch <= z), (ch >= aM && ch <= zM); !min && !may {
			continue
		} else if newCh := ch + rune(k); (min && (newCh > z)) || (may && (newCh > zM)) {
			// Shifts and returns the letter to its own set (a - z) or (A - Z), in case of overflow
			runes[i] = newCh - rune(26)
		} else {
			// Shift the letter
			runes[i] = newCh
		}
	}
	return string(runes)
}
