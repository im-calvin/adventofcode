package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// curr will be all the elements before i
// we multiply curr with nums[i]
func mult(curr int, i int, nums []int) int {
	product := curr * nums[i]

	return product
}

// curr will be all the elements before i
// we add curr with nums[i]
func add(curr int, i int, nums []int) int {
	sum := curr + nums[i]

	return sum
}

func checkAns(a int, b int, result int) int {
	if a == result {
		return a
	} else if b == result {
		return b
	} else {
		return math.MaxInt
	}
}

func helper(curr int, i int, result int, nums []int) int {
	if curr == result {
		return curr
	} else if curr > result {
		return math.MaxInt
	} else if i >= len(nums) {
		return math.MaxInt
	} else {
		product := mult(curr, i, nums)
		sum := add(curr, i, nums)

		res1 := helper(product, i+1, result, nums)
		res2 := helper(sum, i+1, result, nums)

		return checkAns(res1, res2, result)
	}
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)

	calibrationResult := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		lineArr := strings.Split(lineStr, ":")
		result, err := strconv.Atoi(lineArr[0])
		check(err)
		numsArr := strings.Split(strings.TrimSpace(lineArr[1]), " ")
		nums := make([]int, len(numsArr))

		for i, numStr := range numsArr {
			num, err := strconv.Atoi(numStr)
			check(err)
			nums[i] = num
		}

		// start dp
		i := 1
		curr := nums[0]
		res := helper(curr, i, result, nums)

		if res != math.MaxInt {
			calibrationResult = calibrationResult + res
		}

	}

	fmt.Println(calibrationResult)
}
