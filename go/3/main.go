package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var main [256]byte
	var oldCount int = 0

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	for i := 0; i < len(s); i++ {
		var count int = 0
		for k := i; k < len(s); k++ {
			if main[rune(s[k])] > 0 {
				break
			}
			count++
			main[rune(s[k])]++
		}
		if count > oldCount {
			oldCount = count
		}
		main = [256]byte{0}
	}
	return oldCount
}

func main() {
	a := "aab"

	fmt.Println(lengthOfLongestSubstring(a))
}
