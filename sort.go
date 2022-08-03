package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func sort() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	insertionSort(ar)
	bubbleSort(ar)
	selectionSort(ar)
	ar = mergeSort(ar)
	quickSort(ar)

	fmt.Println(ar)
}

//метод сортировки вставкой
func insertionSort(ar []int) {
	key := ar[0]
	for i := range ar {
		key = ar[i]
		for j := i; j >= 0; j-- {
			if key < ar[j] {
				ar[j+1], ar[j] = ar[j], ar[j+1]
			}
		}
	}
}

//метод сортировки пузыриком
func bubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := range ar {
			if j > 0 && ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

//метод сортировки выбором
func selectionSort(ar []int) {
	var minSelected int
	var minIndex int
	length := len(ar)
	for i := range ar {
		minSelected = ar[i]
		minIndex = i
		for j := i; j < length; j++ {
			if ar[j] < minSelected {
				minSelected = ar[j]
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

//метод сортировки слиянием
func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	first := mergeSort(ar[:len(ar)/2])
	second := mergeSort(ar[len(ar)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
