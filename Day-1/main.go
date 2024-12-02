package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	ls1 := []int{3, 4, 2, 1, 3, 3}
	// fmt.Println(ls1)
	sort.Ints(ls1)
	fmt.Println(ls1)

	ls2 := []int{4, 3, 5, 3, 9, 3}
	// fmt.Println(ls2)
	sort.Ints(ls2)
	fmt.Println(ls2)

	// diff1 := abs(ls1[0] - ls2[0])
	// fmt.Println(diff1)

	// diff2 := abs(ls1[1] - ls2[1])
	// fmt.Println(diff2)

	// tot := diff1 + diff2
	// fmt.Println(tot)

	for i:=0;

}
