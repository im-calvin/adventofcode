package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello, World!")
	file, err := os.Open("./input")
	check(err)

	var leftCol []int
	var rightCol []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "   ")
		leftVal, _ := strconv.Atoi(arr[0])
		rightVal, _ := strconv.Atoi(arr[1])
		leftCol = append(leftCol, leftVal)
		rightCol = append(rightCol, rightVal)
	}

	slices.SortFunc(leftCol, func(a, b int) int {
		return a - b
	})

	slices.SortFunc(rightCol, func(a, b int) int {
		return a - b
	})

	if len(leftCol) != len(rightCol) {
		fmt.Printf("Something horrible has happened. Left Col length: %d, Right Col length: %d", len(leftCol), len(rightCol))
	}

	n := len(leftCol)
	var sum int

	for i := range n {
		diff := math.Abs(float64(leftCol[i] - rightCol[i]))
		sum = sum + int(diff)
	}

	fmt.Printf("Answer is %d\n", sum)

	// part 2
	var similarityScore int = 0
	rightFreqMap := make(map[int]int) // maps the number to the frequency

	for _, v := range rightCol {
		rightFreqMap[v] = rightFreqMap[v] + 1
	}

	for _, v := range leftCol {
		similarityScore = similarityScore + v*rightFreqMap[v]
	}

	fmt.Printf("Answer is %d\n", similarityScore)
}
