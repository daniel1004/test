package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"
	result := Check(input)
	fmt.Println(result)
}
func Check(str string) int {
	globalCountrer := 0
	arr := strings.Split(str, "\n")
	for _, v := range arr {
		n1 := strings.Split(v, " ")
		numbers := strings.Split(n1[0], "-")
		min, _ := strconv.Atoi(numbers[0])
		max, _ := strconv.Atoi(numbers[1])
		count := 0
		for i := 0; i < len(n1[2])-1; i++ {
			if n1[2][i] == n1[1][0] {
				count++
			}
		}
		if count >= min && count <= max {
			globalCountrer++
		}
	}
	return globalCountrer

}
