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
	defer file.Close()

	ls1 := make([]string, 0)
	ls2 := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		ls1 = append(ls1, parts[0])
		ls2 = append(ls2, parts[1])
	}

	fmt.Println(ls1[0])
	fmt.Println(ls2[0])

	// ls1 := []int{3, 4, 2, 1, 3, 3}
	// sort.Ints(ls1)

	// ls2 := []int{4, 3, 5, 3, 9, 3}
	// sort.Ints(ls2)

	sum := 0
	for i := 0; i < len(ls1); i++ {
		sum += abs(ls1[i] - ls2[i])
	}
	fmt.Println(sum)
}
