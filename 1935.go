package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	count := 0
	array := strings.Fields(text)
	for _, item := range array {
		broken := false
		for i := 0; i < len(brokenLetters); i++ {
			if strings.Contains(item, string(brokenLetters[i])) {
				broken = true
				break
			}
		}
		if broken == false {
			count = count + 1
		}
	}
	return count
}

func main() {
	count := canBeTypedWords("hello world", "ad")
	//count := canBeTypedWords("leet code", "lt")
	//count := canBeTypedWords("leet code", "e")
	fmt.Println(count)
}
