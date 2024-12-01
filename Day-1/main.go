package main

import (
	"fmt"
	"sort"
)

// var ls1 []int
// var ls2 []int

// ls1 = make([]int 3,4,2,1,3,3)

// func sortInt(xi []int) (ordered []int) {
// 	sort.Ints(xi)
// 	return ordered
// }

func main() {
	ls1 := []int{3, 4, 2, 1, 3, 3}
	// fmt.Println(ls1)
	sort.Ints(ls1)
	fmt.Println(ls1)

	ls2 := []int{4, 3, 5, 3, 9, 3}
	// fmt.Println(ls2)
	sort.Ints(ls2)
	fmt.Println(ls2)

	diff1 := ls1[0] - ls2[0]
	if diff1 < 0 {
		diff1 = -diff1
	}
	fmt.Println(diff1)

	diff2 := ls1[1] - ls2[1]
	if diff2 < 0 {
		diff2 = -diff2
	}
	fmt.Println(diff2)

	tot := diff1 + diff2
	fmt.Println(tot)

	// for _, num1 := range ls1 {
	// 	fmt.Println(num1)
	// }

}
