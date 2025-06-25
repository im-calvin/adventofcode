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

type Direction int

const (
	UP Direction = iota // 0
	RIGHT
	DOWN
	LEFT
)

func (d Direction) Next() Direction {
	return Direction((int(d) + 1) % 4)
}

func forward(x int, y int, direction Direction) (int, int) {
	switch direction {
	case UP:
		x--
	case LEFT:
		y--
	case RIGHT:
		y++
	case DOWN:
		x++
	}

	return x, y
}

func backtrack(x int, y int, direction Direction) (int, int) {

	switch direction {
	case UP:
		x++
	case LEFT:
		y++
	case RIGHT:
		y--
	case DOWN:
		x--
	default:
		direction = UP
	}

	return x, y
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)

	maze := make([][]rune, 0)

	var x, y int

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0)

		for j, char := range line {
			if char == '^' {
				x = len(maze)
				y = j
			}
			row = append(row, char)
		}
		maze = append(maze, row)
	}

	n, m := len(maze), len(maze[0])

	direction := UP // start in the up dir
	// while we are in the bounds
	for x >= 0 && x < n && y >= 0 && y < m {
		// travel to current square
		curr := maze[x][y]
		switch curr {
		case '#':
			// need to backtrack to get off the blockade
			x, y = backtrack(x, y, direction)
			direction = direction.Next() // this will move to the right because of the enum
			continue
		default:
			maze[x][y] = 'X'
		}

		x, y = forward(x, y, direction)
	}

	// count X's
	numVisited := 0
	for _, row := range maze {
		for _, val := range row {
			if val == 'X' {
				numVisited++
			}
		}
	}

	fmt.Println(numVisited)

}
