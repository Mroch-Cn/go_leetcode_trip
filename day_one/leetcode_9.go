package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	target := 0
	var origin = x
	for origin > 0 {
		target = target*10 + origin%10
		origin = origin / 10
	}
	if x == target {
		return true
	}
	return false
}
