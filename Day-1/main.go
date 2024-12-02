package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("/home/trevor/go-projects/ADVENT-OF-CODE-2024/Day-1/numbers.txt")
	if err != nil {
		panic(err)
	}
	// var ls1 []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ls1, ls2, found := strings.Cut(line, " ")
		fmt.Printf("%v\t%v\t%t\n", ls1, ls2, found)
	}

	// ls1 := []int{3, 4, 2, 1, 3, 3}
	// sort.Ints(ls1)

	// ls2 := []int{4, 3, 5, 3, 9, 3}
	// sort.Ints(ls2)

	// sum := 0
	// for i := 0; i < len(ls1); i++ {
	// 	sum += abs(ls1[i] - ls2[i])
	// }
	// fmt.Println(sum)
}
