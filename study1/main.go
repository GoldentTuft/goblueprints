package main

import "fmt"

func main() {
	word := []byte("0123")
	// word = append(word[:2], word[1:]...)
	// fmt.Println(string(word))
	fmt.Println(string(word[:3]) + string(word))
}
