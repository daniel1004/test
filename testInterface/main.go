package main

import "fmt"

func main() {
	x := 123
	isPalindrome(x)
	fmt.Println(x)
}
func isPalindrome(x int) bool {
	dubl := x
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	if reversed != dubl {
		return false
	}
	return true
}
