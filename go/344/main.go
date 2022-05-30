package main

import "fmt"

func main() {
	var s []byte = []byte("hola")
	for i := 0; i < len(s)/2; i++ {
		s[i] += s[len(s)-1-i]
		s[len(s)-1-i] = s[i] - s[len(s)-1-i]
		s[i] -= s[len(s)-1-i]
	}
	fmt.Println(string(s))
}

// a = a + b
// b = a - b
// a = a - b
