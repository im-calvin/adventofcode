package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Coord struct {
	x int
	y int
}

type Queue[T any] struct {
	contents []T
}

func (q Queue[T]) isEmpty() bool {
	return len(q.contents) == 0
}

func (q *Queue[T]) enqueSingle(item T) {
	q.contents = append(q.contents, item)
}

func (q *Queue[T]) enqueSlice(items []T) {
	q.contents = append(q.contents, items...)
}

func (q *Queue[T]) deque() T {
	item := q.contents[0]
	q.contents = slices.Delete(q.contents, 0, 1)
	return item
}

func (q Queue[T]) len() int {
	return len(q.contents)
}

// returns the travel-able nodes
func travel(height int, coord Coord, grid [][]int) []Coord {
	res := make([]Coord, 0)
	if coord.x-1 >= 0 && grid[coord.x-1][coord.y] == height+1 {
		res = append(res, Coord{x: coord.x - 1, y: coord.y})
	}
	if coord.x+1 < len(grid) && grid[coord.x+1][coord.y] == height+1 {
		res = append(res, Coord{x: coord.x + 1, y: coord.y})
	}
	if coord.y-1 >= 0 && grid[coord.x][coord.y-1] == height+1 {
		res = append(res, Coord{x: coord.x, y: coord.y - 1})
	}
	if coord.y+1 < len(grid[0]) && grid[coord.x][coord.y+1] == height+1 {
		res = append(res, Coord{x: coord.x, y: coord.y + 1})
	}
	return res
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)
	grid := make([][]int, 0)
	trailheads := Queue[Coord]{contents: make([]Coord, 0)}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := make([]int, 0)

		for j, char := range line {
			num := int(char - '0')
			nums = append(nums, num)
			if num == 0 {
				trailheads.enqueSingle(Coord{x: i, y: j})
			}
		}
		grid = append(grid, nums)
		i++
	}

	// we should do bfs because then we can check the height positions evenly

	sum := 0
	length := trailheads.len() // the number of nums that still need to be processed
	height := 0

	for !trailheads.isEmpty() {
		coord := trailheads.deque()
		length--
		if grid[coord.x][coord.y] == 9 {
			sum++
		}
		travelable := travel(height, coord, grid)
		trailheads.enqueSlice(travelable)

		if length == 0 {
			length = trailheads.len()
			height++
		}
	}

	fmt.Println(sum)
}
