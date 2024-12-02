package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	ls1 := make([]int, 0)
	ls2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		ls1 = append(ls1, num1)
		ls2 = append(ls2, num2)
	}

	sort.Ints(ls1)
	sort.Ints(ls2)

	sum := 0
	for i := 0; i < len(ls1); i++ {
		sum += abs(ls1[i] - ls2[i])
	}
	fmt.Println(sum)
}
