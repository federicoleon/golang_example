package sort

import (
	"errors"
)

func BubbleSort(numeros []int) error {
	if numeros == nil {
		return errors.New("No se puede procesar un array vacio")
	}
	listo := true
	for listo {
		listo = false
		for i := 0; i < (len(numeros) - 1); i++ {
			if numeros[i] < numeros[i+1] {
				numeros[i], numeros[i+1] = numeros[i+1], numeros[i]
				listo = true
			}
		}
	}
	return nil
}
