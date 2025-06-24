package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// searches the input array at position (x, y) in 8 different ways for the word "SEARCH_TERM"
// left to right, right to left, top to bottom, bottom to top, diagonal top left to bottom right, diagonal top right to bottom left, diagonal bottom left to top right, diagonal bottom right to top left
func search_part_1(input [][]rune, x int, y int, searchTermStr string) int {
	counter := 0 // max possible matches per search
	const WORD_LEN = 4
	searchTerm := []rune(searchTermStr)
	n := len(input)
	m := len(input[0])

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for _, direction := range directions {
		k, l := direction[0], direction[1]

		match := true
		for i := range WORD_LEN {
			dx := x + i*k
			dy := y + i*l

			// check bounds
			if dx < 0 || dx >= n || dy < 0 || dy >= m {
				match = false
				break
			}
			if input[dx][dy] != searchTerm[i] {
				match = false
				break
			}
		}

		if match {
			counter++
		}
	}

	return counter
}

func search_part_2(input [][]rune, x int, y int) int {
	n, m := len(input), len(input[0])

	// search a square around, this time the edges are easier
	if x-1 < 0 || x+1 >= n || y-1 < 0 || y+1 >= m {
		return 0
	}

	// if 'M' is in one corner, then we expect 'S' to be in the other one
	corners := [][]int{
		{-1, -1}, // top left
		{-1, 1},  // top right
	}

	numMatches := 0

	for _, corner := range corners {
		k, l := corner[0], corner[1]

		char := input[x+k][y+l]
		opposite := input[x+k*-1][y+l*-1]

		if (char == 'M' && opposite == 'S') || (char == 'S' && opposite == 'M') {
			numMatches++
		}
	}

	if numMatches == 2 { // we formed a box!
		return 1
	}

	return 0
}

func main() {
	fmt.Println("Hello, World!")
	// const SEARCH_TERM = "XMAS"

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)

	// put the input into a 2d array
	input := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		inner := make([]rune, 0)
		for _, char := range line {
			inner = append(inner, char)
		}
		input = append(input, inner)
	}

	var counter int

	// 2d array for input
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'A' {
				counter = counter + search_part_2(input, i, j)
			}
		}
	}

	fmt.Println(counter)
}
