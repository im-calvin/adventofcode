package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checksum(blocks []Block) int {
	res := 0
	for idx, block := range blocks {
		if block.fileId == -1 {
			continue
		}
		id := block.fileId
		res += id * idx
	}
	return res
}

// swaps 2 nums and returns the stringbuilder swapped
func swap(i int, j int, blocks []Block) []Block {
	blocks[i], blocks[j] = blocks[j], blocks[i]

	return blocks
}

type Block struct {
	fileId int // -1 for free space
}

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)
	var lineStr string
	var disk []Block

	// input is 1 line
	if scanner.Scan() {
		lineStr = scanner.Text()
	}

	id := 0

	for idx, char := range lineStr {
		if idx%2 == 0 {
			// file
			count, err := strconv.Atoi(string(char))
			check(err)
			for range count {
				disk = append(disk, Block{fileId: id})
				check(err)
			}
			id++
		} else {
			// add '.'
			count, err := strconv.Atoi(string(char))
			check(err)
			for range count {
				disk = append(disk, Block{fileId: -1})
				check(err)
			}
		}
	}

	// begin 2 pointer
	head := 0
	tail := len(disk) - 1

	for head < tail {
		if disk[head].fileId == -1 && disk[tail].fileId != -1 {
			disk = swap(head, tail, disk)
			head++
			tail--
		} else if disk[head].fileId != -1 {
			head++
		} else if disk[tail].fileId == -1 {
			tail--
		}
	}

	fmt.Println("Checksum:", checksum(disk))

}
