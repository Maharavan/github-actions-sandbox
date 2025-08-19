package main

import "fmt"

func main() {
	count := []string{"Hello", "Welcome", "Hello", "Greetings"}
	hashMap := make(map[string]int)
	for _, val := range count {
		hashMap[val] += 1
	}
	fmt.Print("Frequency Counter", hashMap)
}
