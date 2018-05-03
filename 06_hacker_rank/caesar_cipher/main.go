package main

import (
	"fmt"
)

func main() {
	var n, k int
	var s string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%d\n", &k)

	fmt.Println(cipherString(s, k))
}

func cipherString(st string, k int) string {
	if k == 0 {
		return st
	}
	runes := []rune(st)
	for i, ch := range runes {
		// Check if the char is not a letter
		if !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') {
			continue
		}
		if newCh := ch + rune(k); (newCh >= 'a' && newCh <= 'z') || (newCh >= 'A' && newCh <= 'Z') {
			runes[i] = newCh
		} else {
			runes[i] = newCh - rune(26)
		}
	}
	return string(runes)
}
