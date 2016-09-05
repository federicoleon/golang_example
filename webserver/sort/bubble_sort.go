package sort

import (
	"errors"
)

func BubbleSort(numbers []int) ([]int, error) {
	if numbers == nil {
		return nil, errors.New("No se puede procesar un array vacio")
	}
	result := make([]int, len(numbers))
	copy(result, numbers)
	listo := true
	for listo {
		listo = false
		for i := 0; i < (len(result) - 1); i++ {
			if result[i] < result[i+1] {
				result[i], result[i+1] = result[i+1], result[i]
				listo = true
			}
		}
	}
	return result, nil
}
