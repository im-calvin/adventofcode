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

func getAntiNodes(new Coordinate, old Coordinate, maxX int, maxY int) []Coordinate {
	res := make([]Coordinate, 0)

	// do negative path
	n := 1
	zCoord := Coordinate{X: old.X - new.X, Y: old.Y - new.Y}
	antiNode := Coordinate{X: old.X - zCoord.X*n, Y: old.Y - zCoord.Y*n}

	for checkNode(antiNode, maxX, maxY) {
		res = append(res, antiNode)
		n++
		antiNode = Coordinate{X: old.X - zCoord.X*n, Y: old.Y - zCoord.Y*n}
	}

	// do positive path
	n = 1
	antiNode = Coordinate{X: new.X + zCoord.X*n, Y: new.Y + zCoord.Y*n}

	for checkNode(antiNode, maxX, maxY) {
		res = append(res, antiNode)
		n++
		antiNode = Coordinate{X: new.X + zCoord.X*n, Y: new.Y + zCoord.Y*n}
	}

	return res
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

		res = append(res, getAntiNodes(new, oldCoord, maxX, maxY)...)
	}

	return res
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./test")
	check(err)

	scanner := bufio.NewScanner(file)
	runeToAntennas := make(map[rune][]Coordinate)

	x := 0
	maxY := 0
	// find all the antennas
	for scanner.Scan() {
		lineStr := scanner.Text()
		maxY = max(maxY, len(lineStr))
		for y, char := range lineStr {
			if char != '.' {
				coord := Coordinate{X: x, Y: y}
				runeToAntennas[char] = append(runeToAntennas[char], coord)
			}
		}
		x++
	}

	maxX := x

	// prepopulate the antiNodeGrid
	antiNodeGrid := make([][]rune, maxX)
	for i := range antiNodeGrid {
		antiNodeGrid[i] = make([]rune, maxY)
		for j := range antiNodeGrid[i] {
			antiNodeGrid[i][j] = '.'
		}
	}

	// find the antiNodes
	for _, antennas := range runeToAntennas {
		for _, antenna := range antennas {
			antinodes := helper(antenna, antennas, maxX, maxY)
			for _, coord := range antinodes {
				antiNodeGrid[coord.X][coord.Y] = '#'
			}
		}
	}

	numAntiNodes := 0

	// count the num of unique antiNodes within the grid
	for _, arr := range antiNodeGrid {
		for _, char := range arr {
			if char == '#' {
				numAntiNodes++
			}
		}
	}

	fmt.Println(numAntiNodes)
}
