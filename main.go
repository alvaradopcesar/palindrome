package main

import (
	"fmt"
	"palindrome/duplicate"
	"palindrome/palindrome"
)

func main() {
	pal := []string{"cesar", "civic", "madam", "radar", "congress", "ala"}
	for _, data := range pal {
		fmt.Println("word :", data, " palindrome : ", palindrome.New().IsPalindrome(data))
	}
	fmt.Println("------")
	result := duplicate.LongStringDuplicate("abcabcdbb")
	fmt.Println(result)

}