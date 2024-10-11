package main

import "fmt"

func repeatedNTimes(nums []int) int {
	l := len(nums) / 2
	mappy := make(map[int]int)

	for _, val := range nums {
		mappy[val]++
		if mappy[val]%l == 0 {
			return val
		}

	}
	return 0
}
func main() {
	a := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(a))

}
