package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) || len(s2) == 0 {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	var main [26]byte
	var tmp [26]byte

	for i := 0; i < len(s1); i++ {
		main[rune(s1[i])-'a']++
		tmp[rune(s2[i])-'a']++
	}

	for i := len(s1); i < len(s2); i++ {
		if tmp == main {
			return true
		}

		tmp[rune(s2[i-len(s1)])-'a']--
		tmp[rune(s2[i])-'a']++
	}
	return tmp == main
}

func main() {
	s1 := "ab"
	s2 := "eeeezxdxasbas"

	fmt.Println(checkInclusion(s1, s2))
}
