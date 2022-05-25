package main

import "fmt"

func main() {
	var s []byte = []byte("Let's start!235")
	var prevSpace int = 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			_len := i - prevSpace
			for k := 0; k < _len/2; k++ {
				s[prevSpace+k] += s[i-1-k]
				s[i-1-k] = s[prevSpace+k] - s[i-1-k]
				s[prevSpace+k] -= s[i-1-k]
			}
			prevSpace = i + 1
		}
	}

	for i := 0; i < (len(s)-prevSpace)/2; i++ {
		s[prevSpace+i] += s[len(s)-1-i]
		s[len(s)-1-i] = s[prevSpace+i] - s[len(s)-1-i]
		s[prevSpace+i] -= s[len(s)-1-i]
	}

	fmt.Println(string(s))

}

// Let's take LeetCode contest
//      5    10       19
