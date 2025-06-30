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

type Visited struct {
	Direction Direction
	X         int
	Y         int
}

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
	// keep track of orientation & location that we traveled to
	traveledNodes := make(map[Visited]bool)
	direction := UP // start in the up dir

	// while we are in the bounds
	for x >= 0 && x < n && y >= 0 && y < m {
		// travel to current square
		curr := maze[x][y]
		visited := Visited{
			direction,
			x,
			y,
		}

		switch curr {
		// golang switch cases won't execute 'X' if '#' executes
		case '#':
			// need to backtrack to get off the blockade
			x, y = backtrack(x, y, direction)
			// if we backtrack we will always land on an 'X'
			direction = direction.Next() // this will move to the right because of the enum
			continue
		default:
			maze[x][y] = 'X'
			traveledNodes[visited] = false // visit the current node with the current direction
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

	numObstructions := 0
	for visited := range traveledNodes {
		// we only want to place obstacles where there haven't already been obstacles!

		// if "false" then the node has been visited but an obstacle hasn't already placed in that spot
		// visited is where I currently am, and obstacle placements happens in front of me. need to check that there is no obstacle in front of me
		x, y := forward(visited.X, visited.Y, visited.Direction)
		if x >= 0 && x < n && y >= 0 && y < m && maze[x][y] == '#' {
			continue
		}

		futureX, futureY := forward(visited.X, visited.Y, visited.Direction.Next())
		futureVisited := Visited{
			visited.Direction.Next(), // the obstacle would reorient, direction.Next() will not mutate `direction`
			futureX,
			futureY,
		}

		if elem, ok := traveledNodes[futureVisited]; ok && !elem {
			numObstructions++
			traveledNodes[visited] = true
		}
	}

	fmt.Println(numVisited)
	fmt.Println(numObstructions)

}
