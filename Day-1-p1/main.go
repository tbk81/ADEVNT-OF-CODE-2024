package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// func that returns the absolute value of an int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// opens the file and closes it after with the defer function
	file, err := os.Open("/home/trevor/go-projects/ADVENT-OF-CODE-2024/Day-1-p1/numbers.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// initialize a slice for each list of numbers
	ls1 := make([]int, 0)
	ls2 := make([]int, 0)

	// scans each line and separates the string in the line into a slice called parts
	// the parts are sub-sliced into num1 and num2, then converted into int's with Atoi
	// each number is appended to the lists initailized above (ls1 and ls2)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		ls1 = append(ls1, num1)
		ls2 = append(ls2, num2)
	}

	// sorts the lists by ascending order
	sort.Ints(ls1)
	sort.Ints(ls2)

	// sums the difference of the numbers in the list and adds them up
	sum := 0
	for i := 0; i < len(ls1); i++ {
		sum += abs(ls1[i] - ls2[i])
	}
	fmt.Println(sum)
}
