package main

import "fmt"

// Нахождение n-го эллемента
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

// Проверка на кратность 3
func spotThree(num int) (int, error) {
	if num%3 == 0 {
		return 0, fmt.Errorf("Число кратно 3 ")
	}
	return num, nil
}

// Input and Output
func main() {
	fmt.Println("Введите количество элементов массива (должно быть кратно 2):")
	var n int
	fmt.Scan(&n)

	// Проверка, что длина кратна 2
	if n%2 != 0 {
		fmt.Println("Ошибка: длина массива должна быть кратна 2.")
		return
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(spotThree(repeatedNTimes(arr)))

}
