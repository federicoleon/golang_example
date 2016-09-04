package sort

import (
	"testing"
)

func TestBubbleSortNoError(t *testing.T) {
	list := []int{6, 3, 1, 5, 7, 0, 2, 0, 8, 4}
	err := BubbleSort(list)
	if err != nil {
		t.Error("we where not expectitn an error")
	}
	for i := 0; i < (len(list) - 1); i++ {
		if list[1] < list[i+1] {
			t.Error("Te slice is or user")
		}
	}
}

func generateSlice(n int) []int {
	var result = make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

func BenchmarkBubbleSortFor10Elements(b *testing.B) {
	list := generateSlice(10)
	for n := 0; n < b.N; n++ {
		BubbleSort(list)
	}
}

func BenchmarkBubbleSortFor100Elements(b *testing.B) {
	list := generateSlice(100)
	for n := 0; n < b.N; n++ {
		BubbleSort(list)
	}
}

func BenchmarkBubbleSortFor1000Elements(b *testing.B) {
	list := generateSlice(1000)
	for n := 0; n < b.N; n++ {
		BubbleSort(list)
	}
}

func BenchmarkBubbleSortFor10000Elements(b *testing.B) {
	list := generateSlice(10000)
	for n := 0; n < b.N; n++ {
		BubbleSort(list)
	}
}
