package sort

func selectionSort(ar []int) {

	for i := 0; i < len(ar)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(ar); j++ {
			if ar[minIndex] > ar[j] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}
func shakerSort(ar []int) {
	n := len(ar)
	left, right := 0, n-1
	for left < right {
		swapped := false
		for i := left; i < right; i++ {
			if ar[i] < ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				swapped = true
			}
		}
		right--
		for i := right; i > left; i-- {
			if ar[i] < ar[i-1] {
				ar[i], ar[i-1] = ar[i-1], ar[i]
				swapped = true
			}
		}
		left++
		if !swapped {
			return
		}
	}
}
func insertionSort(ar []int) {
	n := len(ar)
	for i := 1; i < n; i++ {
		key := ar[i]
		j := i - 1

		for j >= 0 && ar[j] > key {
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1] = key
	}
}
func mergeSort(ar []int) []int {
	//Базовый элемент
	if len(ar) <= 1 {
		return ar
	}
	mid := len(ar) / 2
	//Рекурсивно вызываем сначала left до базового значения, потом right.
	left := mergeSort(ar[:mid])
	right := mergeSort(ar[mid:])
	return merge(left, right)
}
func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
func quickSort(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}
	el := ar[0]
	left := make([]int, 0)
	center := make([]int, 0)
	right := make([]int, 0)
	for _, v := range ar {
		if v < el {
			left = append(left, v)
		} else if v == el {
			center = append(center, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), center...), quickSort(right)...)
}
func bubbleSort(ar []int) []int {
	n := len(ar)
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				swap = true
			}
		}
		if !swap {
			return ar
		}
	}
	return ar
}
