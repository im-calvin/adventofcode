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

type Coordinate struct {
	X int
	Y int
}

func checkNode(coord Coordinate, maxX int, maxY int) bool {
	if coord.X < 0 || coord.X > maxX || coord.Y < 0 || coord.Y > maxY {
		return false
	}
	return true
}

// maxX and maxY are the result of len(grid) and len(grid[0]) - 1, so we need to subtract 1
func helper(new Coordinate, antennas []Coordinate, maxX int, maxY int) []Coordinate {
	res := make([]Coordinate, 0)
	maxX = maxX - 1
	maxY = maxY - 1

	for _, oldCoord := range antennas {
		// antennas can contain "new" so we should leave immediately if that's the case
		if new == oldCoord {
			continue
		}
		zCoord := Coordinate{X: oldCoord.X - new.X, Y: oldCoord.Y - new.Y}
		node1 := Coordinate{X: oldCoord.X - zCoord.X*2, Y: oldCoord.Y - zCoord.Y*2}
		node2 := Coordinate{X: new.X + zCoord.X*2, Y: new.Y + zCoord.Y*2}

		// check that they are in bounds to add them to the result
		if checkNode(node1, maxX, maxY) {
			res = append(res, node1)
		}
		if checkNode(node2, maxX, maxY) {
			res = append(res, node2)
		}
	}

	return res
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)
	runeToAntennas := make(map[rune][]Coordinate)

	x := 0
	maxY := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		maxY = len(lineStr)
		for y, char := range lineStr {
			if char != '.' {
				coord := Coordinate{X: x, Y: y}
				runeToAntennas[char] = append(runeToAntennas[char], coord)
			}
		}
		x++
	}

	maxX := x

	antiNodeGrid := make([][]rune, maxX)
	for i := range antiNodeGrid {
		antiNodeGrid[i] = make([]rune, maxY)
		for j := range antiNodeGrid[i] {
			antiNodeGrid[i][j] = '.'
		}
	}

	for _, antennas := range runeToAntennas {
		for _, antenna := range antennas {
			antinodes := helper(antenna, antennas, maxX, maxY)
			for _, coord := range antinodes {
				antiNodeGrid[coord.X][coord.Y] = '#'
			}
		}
	}

	numAntiNodes := 0

	for _, arr := range antiNodeGrid {
		for _, char := range arr {
			if char == '#' {
				numAntiNodes++
			}
		}
	}

	fmt.Print(numAntiNodes)
}
