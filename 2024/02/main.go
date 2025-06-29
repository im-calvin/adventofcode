package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var numSafe int
	var numFail int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		report := scanner.Text()
		levelStrs := strings.Split(report, " ")
		levels := make([]int, len(levelStrs))

		for i, level := range levelStrs {
			levels[i], _ = strconv.Atoi(level)
		}

		isIncreasing := levels[0] < levels[1]
		damperUsed := false // part 2

		for i, curr := range levels {
			if i == len(levels)-1 {
				numSafe++
				break // avoid indexError
			}
			next := levels[i+1]

			// 1. all levels are either increasing or decreasing
			// if fail then go to next
			// Check if the difference is invalid or direction changes
			if abs(curr, next) == 0 || abs(curr, next) > 3 || (isIncreasing && curr > next) || (!isIncreasing && curr < next) {
				if damperUsed {
					numFail++
					break
				}
				damperUsed = true
			}
		}
	}

	fmt.Println("numSafe:", numSafe)
	fmt.Println("numFail:", numFail)

}
