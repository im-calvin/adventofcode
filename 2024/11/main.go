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

func main() {
	fmt.Println("Hello, World!")

	file, err := os.Open("./input")
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}

	fmt.Println(sum)
}
