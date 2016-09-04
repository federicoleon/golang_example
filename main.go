package main

import (
	"fmt"                                         // Part of https://golang.org/pkg/#stdlib
	"github.com/federicoleon/golang_example/sort" // External dependency
)

func main() {
	var list = []int{1, 4, 6, 6, 8, 4, 2, 4, 5, 7, 89, 9, 4, 2, 23, 6, 7}
	fmt.Println(list)

	sort.BubbleSort(list)

	fmt.Println(list)
}
